package usecase

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"
	"strconv"
	"strings"
	"time"
)

// HELPER FUNCTION
func parseInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 1
}

type ResepJadwalUsecase struct {
	resepRepo      outbound.ResepObatRepository
	jadwalObatRepo outbound.JadwalObatRepository
	jadwalRepo     outbound.JadwalRepository
	obatRepo       outbound.ObatRepository
}

func NewResepJadwalUsecase(r outbound.ResepObatRepository, j outbound.JadwalObatRepository, jr outbound.JadwalRepository, o outbound.ObatRepository) *ResepJadwalUsecase {
	return &ResepJadwalUsecase{
		resepRepo:      r,
		jadwalObatRepo: j,
		jadwalRepo:     jr,
		obatRepo:       o,
	}
}

func generateJamMinum(frekuensi int) []string {
	switch frekuensi {
	case 1:
		return []string{"08:00"}
	case 2:
		return []string{"08:00", "20:00"}
	case 3:
		return []string{"08:00", "14:00", "20:00"}
	default:
		return []string{"08:00"}
	}
}

func (u *ResepJadwalUsecase) Create(req *dto.CreateResepWithJadwalDTO) error {

	// parseDate: support format YYYY-MM-DD dan RFC3339
	parseDate := func(s string) (time.Time, error) {
		if t, err := time.Parse("2006-01-02", s); err == nil {
			return t, nil
		}
		return time.Parse(time.RFC3339, s)
	}

	// PARSE TANGGAL MULAI
	tanggalMulai, err := parseDate(req.TanggalMulai)
	if err != nil {
		return err
	}

	// HANDLE JAM MINUM
	jamList := req.JamMinum
	if len(jamList) == 0 {
		jamList = generateJamMinum(req.FrekuensiPerHari)
	}

	// Cek jika mulai hari ini dan semua jam minum sudah terlewat, geser ke besok
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if loc == nil {
		loc = time.Local
	}
	todayStr := time.Now().In(loc).Format("2006-01-02")
	if tanggalMulai.Format("2006-01-02") == todayStr {
		allPassed := true
		now := time.Now().In(loc)
		currentHour, currentMin, _ := now.Clock()
		for _, jam := range jamList {
			parts := strings.Split(strings.TrimSpace(jam), ":")
			if len(parts) == 2 {
				hour, err1 := strconv.Atoi(parts[0])
				min, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					if hour > currentHour || (hour == currentHour && min >= currentMin) {
						allPassed = false
						break
					}
				}
			}
		}
		if allPassed && len(jamList) > 0 {
			tanggalMulai = tanggalMulai.AddDate(0, 0, 1)
		}
	}

	// Handle tanggal_selesai
	var tanggalSelesai time.Time
	if req.TipeDurasi == "hari" && req.JumlahHari > 0 {
		tanggalSelesai = tanggalMulai.AddDate(0, 0, req.JumlahHari - 1)
	} else if req.TanggalSelesai != "" && req.TanggalSelesai != "null" {
		tanggalSelesai, err = parseDate(req.TanggalSelesai)
		if err != nil {
			return err
		}
	} else if req.TipeDurasi == "rutin" {
		// Kosongkan tanggal selesai untuk rutin
	} else {
		// Jika kosong dan bukan rutin, set default ke 30 hari
		tanggalSelesai = tanggalMulai.AddDate(0, 0, 30)
	}

	// ITERASI OBAT LIST UNTUK MEMBUAT RESEP DAN JADWAL OBAT MASING-MASING
	var combinedNamaObatList []string
	var combinedTakaranList []string

	for _, obatItem := range req.ObatList {
		// AMBIL DATA OBAT DARI MASTER OBAT
		obat, err := u.obatRepo.GetByID(obatItem.ObatID)
		if err != nil {
			return err
		}
		
		// Kumpulkan nama dan takaran untuk Jadwal gabungan
		namaDenganDosis := obat.NamaObat + " (" + obatItem.Dosis + ")"
		combinedNamaObatList = append(combinedNamaObatList, namaDenganDosis)
		combinedTakaranList = append(combinedTakaranList, obatItem.Dosis)

		// BUAT RESEP
		resep := &domain.ResepObat{
			PasienID: req.PasienID,
			ObatID:   obatItem.ObatID,
			NakesID:  req.NakesID,
			Dosis:    obatItem.Dosis,
			Catatan:  req.Catatan,

			TanggalMulai:   tanggalMulai,
			TanggalSelesai: tanggalSelesai,

			FrekuensiPerHari: req.FrekuensiPerHari,
			AturanKonsumsi:   req.AturanKonsumsi,
		}

		// SIMPAN RESEP
		if err := u.resepRepo.Create(resep); err != nil {
			return err
		}

		// SIMPAN JADWAL OBAT
		for _, jam := range jamList {
			parsedTime, err := time.Parse("15:04", jam)
			if err != nil {
				return err
			}

			jadwalObat := &domain.JadwalObat{
				ResepObatID: resep.ResepObatID,
				PasienID:    req.PasienID,
				ObatID:      obatItem.ObatID,
				NakesID:     req.NakesID,
				JamMinum:    parsedTime,
			}

			if err := u.jadwalObatRepo.Create(jadwalObat); err != nil {
				return err
			}
		}
	}

	// BUAT 1 JADWAL DI TABEL JADWAL (UNTUK JADWAL-SERVICE)
	// Hitung jumlah hari
	jumlahHari := req.JumlahHari
	if jumlahHari == 0 && !tanggalSelesai.IsZero() && !tanggalMulai.IsZero() {
		jumlahHari = int(tanggalSelesai.Sub(tanggalMulai).Hours() / 24)
	}

	tipeDurasi := req.TipeDurasi
	if tipeDurasi == "" {
		tipeDurasi = "hari"
	}

	// Format tanggal ke string format yang diharapkan tabel jadwal
	tglMulaiStr := tanggalMulai.Format("2006-01-02")
	tglSelesaiStr := ""
	if !tanggalSelesai.IsZero() {
		tglSelesaiStr = tanggalSelesai.Format("2006-01-02")
	}

	// Gabungkan nama obat dan takaran
	namaObatGabungan := strings.Join(combinedNamaObatList, ", ")
	takaranGabungan := strings.Join(combinedTakaranList, ", ")

	// Buat jadwal entry di tabel jadwal
	jadwalForService := &domain.Jadwal{
		PasienID:           req.PasienID,
		NamaObat:           namaObatGabungan,
		JumlahDosis:        1,        // Placeholder, detail di NamaObat
		Satuan:             "custom", // Placeholder
		KategoriObat:       "multi",
		TakaranObat:        takaranGabungan,
		FrekuensiPerHari:   strconv.Itoa(req.FrekuensiPerHari),
		WaktuMinum:         strings.Join(jamList, ", "),
		AturanKonsumsi:     req.AturanKonsumsi,
		Catatan:            req.Catatan,
		TipeDurasi:         tipeDurasi,
		JumlahHari:         jumlahHari,
		TanggalMulai:       tglMulaiStr,
		TanggalSelesai:     tglSelesaiStr,
		Status:             "aktif",
	}

	_, err = u.jadwalRepo.Create(jadwalForService)
	if err != nil {
		return err
	}

	return nil
}
