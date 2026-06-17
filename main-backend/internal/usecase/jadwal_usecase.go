package usecase

import (
	"backend/internal/adapters/outbound/persistence"
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"
	"backend/pkg/fcm"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type JadwalUsecase struct {
	jadwalRepo   outbound.JadwalRepository
	pasienRepo   outbound.PasienRepository
	fcmTokenRepo outbound.FcmTokenRepository
}

func NewJadwalUsecase(jr outbound.JadwalRepository, pr outbound.PasienRepository, fcmRepo outbound.FcmTokenRepository) *JadwalUsecase {
	return &JadwalUsecase{
		jadwalRepo:   jr,
		pasienRepo:   pr,
		fcmTokenRepo: fcmRepo,
	}
}

func (u *JadwalUsecase) GetAllJadwal() ([]dto.JadwalResponseDTO, error) {
	jadwals, err := u.jadwalRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.JadwalResponseDTO
	for _, jadwal := range jadwals {
		dto := persistence.JadwalToResponseDTO(&jadwal)
		responses = append(responses, *dto)
	}

	return responses, nil
}

func (u *JadwalUsecase) GetJadwalByID(id int) (*dto.JadwalResponseDTO, error) {
	jadwal, err := u.jadwalRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return persistence.JadwalToResponseDTO(jadwal), nil
}

func (u *JadwalUsecase) GetJadwalByPasien(pasienID int) ([]dto.JadwalResponseDTO, error) {
	jadwals, err := u.jadwalRepo.GetByPasienID(pasienID)
	if err != nil {
		return nil, err
	}

	var responses []dto.JadwalResponseDTO
	for _, jadwal := range jadwals {
		dto := persistence.JadwalToResponseDTO(&jadwal)
		responses = append(responses, *dto)
	}

	return responses, nil
}

func (u *JadwalUsecase) CreateJadwal(req *dto.CreateJadwalDTO) (*dto.JadwalResponseDTO, error) {
	if req.PasienID <= 0 {
		return nil, errors.New("pasien_id harus valid")
	}

	// Validate pasien exists
	pasien, _ := u.pasienRepo.GetByID(uint(req.PasienID))
	if pasien == nil || pasien.PasienID == 0 {
		return nil, errors.New("pasien tidak ditemukan")
	}

	// Hitung tanggal selesai jika tipe_durasi = "hari"
	tanggalMulaiStr := req.TanggalMulai
	tanggalSelesai := req.TanggalSelesai
	if req.TanggalMulai != "" {
		mulai, err := time.Parse("2006-01-02", req.TanggalMulai)
		if err == nil {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			if loc == nil {
				loc = time.Local
			}
			todayStr := time.Now().In(loc).Format("2006-01-02")
			if mulai.Format("2006-01-02") == todayStr {
				// cek jika semua jam minum sudah terlewat hari ini
				allPassed := true
				now := time.Now().In(loc)
				currentHour, currentMin, _ := now.Clock()
				timesList := strings.Split(req.WaktuMinum, ",")
				for _, jam := range timesList {
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
				if allPassed && len(timesList) > 0 && req.WaktuMinum != "" {
					mulai = mulai.AddDate(0, 0, 1)
					tanggalMulaiStr = mulai.Format("2006-01-02")
				}
			}
			if req.TipeDurasi == "hari" && req.JumlahHari > 0 {
				selesai := mulai.AddDate(0, 0, req.JumlahHari - 1)
				tanggalSelesai = selesai.Format("2006-01-02")
			}
		}
	}

	jadwal := &domain.Jadwal{
		PasienID:           req.PasienID,
		NamaObat:           req.NamaObat,
		JumlahDosis:        req.JumlahDosis,
		Satuan:             req.Satuan,
		KategoriObat:       req.KategoriObat,
		TakaranObat:        req.TakaranObat,
		FrekuensiPerHari:   req.FrekuensiPerHari,
		WaktuMinum:         req.WaktuMinum,
		AturanKonsumsi:     req.AturanKonsumsi,
		Catatan:            req.Catatan,
		TipeDurasi:         req.TipeDurasi,
		JumlahHari:         req.JumlahHari,
		TanggalMulai:       tanggalMulaiStr,
		TanggalSelesai:     tanggalSelesai,
		Status:             "aktif",
	}

	result, err := u.jadwalRepo.Create(jadwal)
	if err != nil {
		return nil, err
	}

	// Isi PasienNama secara manual karena Create tidak JOIN
	result.PasienNama = pasien.Nama
	responseDTO := persistence.JadwalToResponseDTO(result)

	// ── Kirim FCM ke device pasien (non-blocking) ──
	// Error FCM tidak menggagalkan create jadwal
	go func() {
		token, err := u.fcmTokenRepo.GetByPasienID(uint(req.PasienID))
		if err != nil || token == "" {
			log.Printf("[FCM] Token tidak ditemukan untuk pasien %d: %v", req.PasienID, err)
			return
		}

		payload := map[string]string{
			"type":                 "jadwal_baru",
			"jadwal_id":            fmt.Sprintf("%d", result.JadwalID),
			"nama_obat":            result.NamaObat,
			"jumlah_dosis":         fmt.Sprintf("%d", result.JumlahDosis),
			"satuan":               result.Satuan,
			"waktu_minum":          result.WaktuMinum,
			"kategori_obat":        result.KategoriObat,
			"aturan_konsumsi":      result.AturanKonsumsi,
			"tanggal_mulai":        result.TanggalMulai,
			"tanggal_selesai":      result.TanggalSelesai,
			"status":               result.Status,
			"title":                "💊 Jadwal Obat Baru",
			"body":                 fmt.Sprintf("Nakes menambahkan jadwal %s pukul %s", result.NamaObat, result.WaktuMinum),
		}

		if err := fcm.SendJadwalNotification(context.Background(), token, payload); err != nil {
			log.Printf("[FCM] Gagal kirim notifikasi jadwal ke pasien %d: %v", req.PasienID, err)
		}
	}()

	return responseDTO, nil
}

func (u *JadwalUsecase) UpdateJadwal(id int, req *dto.UpdateJadwalDTO) (*dto.JadwalResponseDTO, error) {
	jadwal, err := u.jadwalRepo.GetByID(id)
	if err != nil || jadwal == nil {
		return nil, errors.New("jadwal tidak ditemukan")
	}

	// Update fields yang tidak kosong
	if req.NamaObat != "" {
		jadwal.NamaObat = req.NamaObat
	}
	if req.JumlahDosis > 0 {
		jadwal.JumlahDosis = req.JumlahDosis
	}
	if req.Satuan != "" {
		jadwal.Satuan = req.Satuan
	}
	if req.KategoriObat != "" {
		jadwal.KategoriObat = req.KategoriObat
	}
	if req.TakaranObat != "" {
		jadwal.TakaranObat = req.TakaranObat
	}
	if req.FrekuensiPerHari != "" {
		jadwal.FrekuensiPerHari = req.FrekuensiPerHari
	}
	if req.WaktuMinum != "" {
		jadwal.WaktuMinum = req.WaktuMinum
	}
	if req.AturanKonsumsi != "" {
		jadwal.AturanKonsumsi = req.AturanKonsumsi
	}
	if req.Catatan != "" {
		jadwal.Catatan = req.Catatan
	}
	if req.TipeDurasi != "" {
		jadwal.TipeDurasi = req.TipeDurasi
	}
	if req.JumlahHari > 0 {
		jadwal.JumlahHari = req.JumlahHari
	}
	if req.TanggalMulai != "" {
		jadwal.TanggalMulai = req.TanggalMulai
	}
	if req.TanggalSelesai != "" {
		jadwal.TanggalSelesai = req.TanggalSelesai
	}
	if req.Status != "" {
		jadwal.Status = req.Status
	}

	// Update dan re-fetch dengan JOIN agar PasienNama terisi
	result, err := u.jadwalRepo.Update(id, jadwal)
	if err != nil {
		return nil, err
	}

	return persistence.JadwalToResponseDTO(result), nil
}

func (u *JadwalUsecase) DeleteJadwal(id int) error {
	return u.jadwalRepo.Delete(id)
}
