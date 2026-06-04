package usecase

import (
	"backend/internal/adapters/outbound/persistence"
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ObatUsecase struct {
	repo       outbound.ObatRepository
	jadwalRepo outbound.JadwalRepository
}

func NewObatUsecase(r outbound.ObatRepository, j outbound.JadwalRepository) *ObatUsecase {
	return &ObatUsecase{repo: r, jadwalRepo: j}
}

// GetAll mendapatkan semua obat sesuai kolom DB (hanya obat milik admin, menyaring obat mandiri)
func (u *ObatUsecase) GetAll() ([]dto.ObatResponseDTO, error) {
	obats, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := []dto.ObatResponseDTO{}
	for _, obat := range obats {
		if !obat.IsMandiri {
			responses = append(responses, *persistence.ObatToResponseDTO(&obat))
		}
	}
	return responses, nil
}

// GetAllByPasien mendapatkan semua obat mandiri milik pasien tertentu
func (u *ObatUsecase) GetAllByPasien(pasienID int) ([]dto.ObatResponseDTO, error) {
	obats, err := u.repo.GetByPasienID(pasienID)
	if err != nil {
		return nil, err
	}

	responses := []dto.ObatResponseDTO{}
	for _, obat := range obats {
		responses = append(responses, *persistence.ObatToResponseDTO(&obat))
	}
	return responses, nil
}

// GetByID mendapatkan obat berdasarkan ID
func (u *ObatUsecase) GetByID(id int) (*dto.ObatResponseDTO, error) {
	obat, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if obat == nil || obat.ObatID == 0 {
		return nil, errors.New("obat tidak ditemukan")
	}
	return persistence.ObatToResponseDTO(obat), nil
}

// Create membuat obat baru (Hanya menggunakan kolom yang ada di DB)
func (u *ObatUsecase) Create(req *dto.CreateObatDTO) (*dto.ObatResponseDTO, error) {
	// 1. Validasi Input Dasar
	if req.NamaObat == "" {
		return nil, errors.New("nama obat tidak boleh kosong")
	}

	// 2. Cek Duplikasi Nama
	existing, _ := u.repo.GetByName(req.NamaObat)
	if existing != nil && existing.ObatID != 0 {
		return nil, errors.New("obat dengan nama tersebut sudah ada")
	}

	// 3. Mapping DTO ke Domain (Hanya kolom SQL: nama, fungsi, aturan_pemakaian, pantangan, gambar)
	obat := &domain.Obat{
		NamaObat:        req.NamaObat,
		Fungsi:          req.Fungsi,
		AturanPemakaian: req.AturanPemakaian, // Sesuai kolom DB
		Pantangan:       req.Pantangan,       // Sesuai kolom DB
		Gambar:          req.Gambar,          // Path dari handler
	}

	// 4. Simpan ke Persistence
	result, err := u.repo.Create(obat)
	if err != nil {
		return nil, errors.New("gagal menyimpan data obat ke database")
	}

	return persistence.ObatToResponseDTO(result), nil
}

// Update memperbarui data obat (Sesuai kolom DB)
func (u *ObatUsecase) Update(id int, req *dto.UpdateObatDTO) (*dto.ObatResponseDTO, error) {
	// 1. Cari data lama
	existing, err := u.repo.GetByID(id)
	if err != nil || existing == nil || existing.ObatID == 0 {
		return nil, errors.New("obat tidak ditemukan")
	}

	oldNamaObat := existing.NamaObat

	// 2. Partial Update (Hanya kolom yang ada di SQL)
	if req.NamaObat != "" {
		existing.NamaObat = req.NamaObat
	}
	if req.Fungsi != "" {
		existing.Fungsi = req.Fungsi
	}
	if req.AturanPemakaian != "" {
		existing.AturanPemakaian = req.AturanPemakaian
	}
	if req.Pantangan != "" {
		existing.Pantangan = req.Pantangan
	}

	// 3. Logika Gambar: Tetap dipertahankan
	if req.Gambar != "" {
		existing.Gambar = req.Gambar
	}

	// 3.5 Kolom Tambahan untuk Obat Mandiri Pasien
	if existing.IsMandiri {
		if req.Dosis != "" {
			existing.Fungsi = req.Dosis
		}
		if req.Frekuensi != "" {
			existing.Frekuensi = req.Frekuensi
		}
		if req.DurasiHari > 0 {
			existing.DurasiHari = &req.DurasiHari
		}
		if req.Catatan != "" {
			existing.Catatan = req.Catatan
		}
		if len(req.Pengingat) > 0 {
			_ = existing.SetPengingat(req.Pengingat)
		}
	}

	// 4. Eksekusi Update di Repo
	result, err := u.repo.Update(id, existing)
	if err != nil {
		return nil, errors.New("gagal memperbarui data obat")
	}

	// 5. Update jadwal terkait jika ini obat mandiri
	if existing.IsMandiri && existing.PasienID != nil {
		jadwals, err := u.jadwalRepo.GetByPasienID(*existing.PasienID)
		if err == nil {
			for _, j := range jadwals {
				if j.NamaObat == oldNamaObat && j.KategoriObat == "Mandiri" {
					j.NamaObat = existing.NamaObat
					j.TakaranObat = existing.Fungsi // Dosis
					j.AturanKonsumsi = existing.Catatan
					j.Catatan = existing.Catatan
					if existing.DurasiHari != nil {
						j.JumlahHari = *existing.DurasiHari
					}
					
					var times []string
					if len(req.Pengingat) > 0 {
						times = req.Pengingat
					} else {
						times = existing.GetPengingat()
					}
					sort.Strings(times)
					j.WaktuMinum = strings.Join(times, ", ")
					j.FrekuensiPerHari = strconv.Itoa(len(times))
					
					if existing.DurasiHari != nil && *existing.DurasiHari > 0 {
						if tMulai, err := time.Parse("2006-01-02", j.TanggalMulai); err == nil {
							selesai := tMulai.AddDate(0, 0, *existing.DurasiHari - 1)
							j.TanggalSelesai = selesai.Format("2006-01-02")
						}
					}
					
					_, _ = u.jadwalRepo.Update(j.JadwalID, &j)
				}
			}
		}
	}

	return persistence.ObatToResponseDTO(result), nil
}

// Delete menghapus obat
func (u *ObatUsecase) Delete(id int) error {
	existing, err := u.repo.GetByID(id)
	if err != nil || existing == nil || existing.ObatID == 0 {
		return errors.New("obat tidak ditemukan")
	}

	// Hapus jadwal terkait jika ini obat mandiri
	if existing.IsMandiri && existing.PasienID != nil {
		jadwals, err := u.jadwalRepo.GetByPasienID(*existing.PasienID)
		if err == nil {
			for _, j := range jadwals {
				if j.NamaObat == existing.NamaObat && j.KategoriObat == "Mandiri" {
					_ = u.jadwalRepo.Delete(j.JadwalID)
				}
			}
		}
	}

	return u.repo.Delete(id)
}

// CreateMandiri membuat obat mandiri untuk pasien
func (u *ObatUsecase) CreateMandiri(req *dto.CreateObatMandiriDTO) (*dto.ObatResponseDTO, error) {
	// 1. Validasi Input
	if req.NamaObat == "" {
		return nil, errors.New("nama obat tidak boleh kosong")
	}
	if req.Dosis == "" {
		return nil, errors.New("dosis tidak boleh kosong")
	}
	if len(req.Pengingat) == 0 {
		return nil, errors.New("pengingat tidak boleh kosong")
	}
	if req.Frekuensi == "" {
		return nil, errors.New("frekuensi tidak boleh kosong")
	}
	if req.DurasiHari <= 0 {
		return nil, errors.New("durasi hari harus lebih dari 0")
	}
	if req.PasienID <= 0 {
		return nil, errors.New("pasien_id tidak valid")
	}

	// 2. Mapping DTO ke Domain
	obat := &domain.Obat{
		NamaObat:   req.NamaObat,
		Gambar:     req.Gambar,
		PasienID:   &req.PasienID,
		Frekuensi:  req.Frekuensi,
		DurasiHari: &req.DurasiHari,
		Catatan:    req.Catatan,
		IsMandiri:  true,
	}

	// Set pengingat sebagai JSON array
	if err := obat.SetPengingat(req.Pengingat); err != nil {
		return nil, errors.New("gagal menyimpan waktu pengingat")
	}

	// Dosis diisi di field Fungsi untuk kompatibilitas
	obat.Fungsi = req.Dosis

	// 3. Simpan ke Database
	result, err := u.repo.Create(obat)
	if err != nil {
		return nil, errors.New("gagal menyimpan data obat mandiri: " + err.Error())
	}

	// 4. Buat Jadwal otomatis di tabel jadwal
	// Urutkan pengingat secara kronologis
	sort.Strings(req.Pengingat)
	waktuMinum := strings.Join(req.Pengingat, ", ")

	loc, _ := time.LoadLocation("Asia/Jakarta")
	if loc == nil {
		loc = time.Local
	}
	now := time.Now().In(loc)
	tanggalMulai := now

	// Cek jika seluruh pengingat sudah lewat hari ini, geser mulai ke besok
	allPassed := true
	currentHour, currentMin, _ := now.Clock()
	for _, jam := range req.Pengingat {
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
	if allPassed && len(req.Pengingat) > 0 {
		tanggalMulai = tanggalMulai.AddDate(0, 0, 1)
	}

	todayStr := tanggalMulai.Format("2006-01-02")
	tanggalSelesaiStr := ""
	if req.DurasiHari > 0 {
		selesai := tanggalMulai.AddDate(0, 0, req.DurasiHari - 1)
		tanggalSelesaiStr = selesai.Format("2006-01-02")
	}

	jadwal := &domain.Jadwal{
		PasienID:           req.PasienID,
		NamaObat:           req.NamaObat,
		JumlahDosis:        1,
		Satuan:             "dosis",
		KategoriObat:       "Mandiri",
		TakaranObat:        req.Dosis,
		FrekuensiPerHari:   strconv.Itoa(len(req.Pengingat)),
		WaktuMinum:         waktuMinum,
		AturanKonsumsi:     req.Catatan,
		Catatan:            req.Catatan,
		TipeDurasi:         "hari",
		JumlahHari:         req.DurasiHari,
		TanggalMulai:       todayStr,
		TanggalSelesai:     tanggalSelesaiStr,
		WaktuReminderPagi:  "",
		WaktuReminderMalam: "",
		Status:             "aktif",
	}

	_, err = u.jadwalRepo.Create(jadwal)
	if err != nil {
		return nil, errors.New("gagal menyimpan jadwal otomatis untuk obat mandiri: " + err.Error())
	}

	return persistence.ObatToResponseDTO(result), nil
}