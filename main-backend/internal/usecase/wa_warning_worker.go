package usecase

import (
	"backend/internal/domain"
	"backend/pkg/whatsapp"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type WaWarningWorker struct {
	db *gorm.DB
}

func NewWaWarningWorker(db *gorm.DB) *WaWarningWorker {
	return &WaWarningWorker{db: db}
}

func (w *WaWarningWorker) Start() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		log.Println("[WA Worker] Pekerja Latar Belakang Peringatan WhatsApp dimulai")
		for range ticker.C {
			w.CheckAndSendWarnings()
		}
	}()
}

func (w *WaWarningWorker) CheckAndSendWarnings() {
	// Menggunakan FixedZone agar tidak terjadi error timezone di Windows
	loc := time.FixedZone("WIB", 7*3600)

	now := time.Now().In(loc)
	todayStr := now.Format("2006-01-02")
	
	// Parsing tanggal hari ini ke time.Time untuk pembandingan di DB GORM
	todayDate, parseErr := time.ParseInLocation("2006-01-02", todayStr, loc)
	if parseErr != nil {
		log.Println("[WA Worker] Gagal melakukan parse tanggal hari ini:", parseErr)
		return
	}

	// Ambil semua jadwal obat yang berstatus 'aktif' atau 'active' (termasuk obat mandiri)
	var jadwals []domain.Jadwal
	if err := w.db.Where("status = ? OR status = ?", "aktif", "active").Find(&jadwals).Error; err != nil {
		log.Println("[WA Worker] Gagal mengambil daftar jadwal aktif dari database:", err)
		return
	}

	for _, j := range jadwals {
		// Validasi apakah jadwal berlaku untuk hari ini
		if j.TanggalMulai > todayStr {
			continue
		}
		if j.TanggalSelesai != "" && j.TanggalSelesai < todayStr {
			continue
		}

		// Split waktu_minum jika jadwal memiliki beberapa jam minum terpisah koma
		waktuList := []string{j.WaktuMinum}
		if strings.Contains(j.WaktuMinum, ",") {
			parts := strings.Split(j.WaktuMinum, ",")
			waktuList = []string{}
			for _, part := range parts {
				waktuList = append(waktuList, strings.TrimSpace(part))
			}
		}

		for _, wm := range waktuList {
			if wm == "" {
				continue
			}

			// Parse waktu jadwal hari ini lengkap
			schedTimeFull, err := time.ParseInLocation("2006-01-02 15:04", todayStr+" "+wm, loc)
			if err != nil {
				continue
			}

			diff := now.Sub(schedTimeFull)
			elapsedMinutes := int(diff.Minutes())

			// DEBUG LOG: Agar kita bisa melihat di terminal status pengecekan tiap menit
			// log.Printf("[WA Worker] Cek JadwalID %d | Waktu Jadwal: %s | Sudah lewat: %d menit", j.JadwalID, wm, elapsedMinutes)

			// Peringatan dikirimkan ketika tersisa 20 menit dari batas waktu 60 menit (artinya di menit ke-40)
			// Kita menggunakan rentang [40, 60] menit.
			if elapsedMinutes >= 40 && elapsedMinutes < 60 {
				
				// 1. Periksa apakah pasien sudah meminum obat tersebut hari ini
				var complianceCount int64
				err := w.db.Model(&domain.TrackingJadwal{}).
					Where("jadwal_id = ? AND tanggal = ? AND status IN (?, ?)", j.JadwalID, todayDate, "Diminum", "Terlambat").
					Count(&complianceCount).Error
				if err != nil {
					log.Printf("[WA Worker] Gagal memeriksa kepatuhan untuk JadwalID %d: %v", j.JadwalID, err)
					continue
				}

				if complianceCount > 0 {
					// Pasien sudah meminum/mengonfirmasi dosis obat ini
					continue
				}

				// 2. Periksa apakah peringatan WA sudah pernah dikirimkan hari ini untuk jadwal tersebut
				var warningCount int64
				err = w.db.Model(&domain.WaWarning{}).
					Where("jadwal_id = ? AND tanggal = ?", j.JadwalID, todayDate).
					Count(&warningCount).Error
				if err != nil {
					log.Printf("[WA Worker] Gagal memeriksa status log peringatan untuk JadwalID %d: %v", j.JadwalID, err)
					continue
				}

				if warningCount > 0 {
					// Peringatan sudah dikirim hari ini
					continue
				}

				// 3. Ambil data profil pasien untuk membaca no_telepon_pendamping
				var pasien domain.Pasien
				if err := w.db.First(&pasien, j.PasienID).Error; err != nil {
					log.Printf("[WA Worker] Gagal memuat profil pasien ID %d: %v", j.PasienID, err)
					continue
				}

				// Jika no_telepon_pendamping terisi, kirim peringatan
				if pasien.NoTeleponPendamping != "" {
					// Panggil simulator pengiriman WhatsApp
					whatsapp.SendWarning(
						pasien.NoTeleponPendamping,
						pasien.Nama,
						j.NamaObat,
						wm,
					)

					// 4. Catat log pengiriman peringatan ke tabel wa_warnings
					warningLog := domain.WaWarning{
						JadwalID: j.JadwalID,
						Tanggal:  todayDate,
						SentAt:   now,
					}
					if err := w.db.Create(&warningLog).Error; err != nil {
						log.Printf("[WA Worker] Gagal menyimpan log pengiriman peringatan ke DB: %v", err)
					}
				} else {
					log.Printf("[WA Worker] Pasien '%s' belum menentukan nomor WA Pendamping. Peringatan tidak dapat dikirim.", pasien.Nama)
				}
			}
		}
	}
}
