package usecase

import (
	"backend/internal/adapters/outbound/persistence"
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"
	"backend/internal/utils"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// ==================== TrackingJadwal Usecase ====================

type TrackingJadwalUsecase struct {
	trackingRepo outbound.TrackingJadwalRepository
	jadwalRepo   outbound.JadwalRepository
	pasienRepo   outbound.PasienRepository
}

func NewTrackingJadwalUsecase(
	tr outbound.TrackingJadwalRepository,
	jr outbound.JadwalRepository,
	pr outbound.PasienRepository,
) *TrackingJadwalUsecase {
	return &TrackingJadwalUsecase{
		trackingRepo: tr,
		jadwalRepo:   jr,
		pasienRepo:   pr,
	}
}

func (u *TrackingJadwalUsecase) GetAll() ([]dto.TrackingJadwalDTO, error) {
	trackings, err := u.trackingRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.TrackingJadwalDTO
	for _, tracking := range trackings {
		jadwal, _ := u.jadwalRepo.GetByID(tracking.JadwalID)
		namaObat := ""
		if jadwal != nil {
			if jadwal.KategoriObat == "Mandiri" {
				continue
			}
			namaObat = jadwal.NamaObat
		}

		pasien, _ := u.pasienRepo.GetByID(uint(tracking.PasienID))
		namaPasien := ""
		if pasien != nil {
			namaPasien = pasien.Nama
		}

		dto := persistence.TrackingJadwalToDTO(&tracking, namaObat, namaPasien)
		responses = append(responses, *dto)
	}

	// ── GENERATE DYNAMIC MISSED TRACKINGS FOR ALL ACTIVE PATIENT SCHEDULES ──
	jadwals, err := u.jadwalRepo.GetAll()
	if err == nil {
		trackingsByPasien := make(map[int][]domain.TrackingJadwal)
		for _, t := range trackings {
			trackingsByPasien[t.PasienID] = append(trackingsByPasien[t.PasienID], t)
		}

		jadwalsByPasien := make(map[int][]domain.Jadwal)
		for _, j := range jadwals {
			if j.KategoriObat != "Mandiri" {
				jadwalsByPasien[j.PasienID] = append(jadwalsByPasien[j.PasienID], j)
			}
		}

		for pID, pJadwals := range jadwalsByPasien {
			pasien, _ := u.pasienRepo.GetByID(uint(pID))
			namaPasien := ""
			if pasien != nil {
				namaPasien = pasien.Nama
			}
			pTrackings := trackingsByPasien[pID]
			missedResponses := u.generateDynamicMissedTrackings(pTrackings, pJadwals, namaPasien)
			responses = append(responses, missedResponses...)
		}
	}

	// Sort responses by Tanggal descending
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Tanggal > responses[j].Tanggal
	})

	return responses, nil
}

func (u *TrackingJadwalUsecase) GetByID(id int) (*dto.TrackingJadwalDTO, error) {
	tracking, err := u.trackingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	jadwal, _ := u.jadwalRepo.GetByID(tracking.JadwalID)
	namaObat := ""
	if jadwal != nil {
		namaObat = jadwal.NamaObat
	}

	pasien, _ := u.pasienRepo.GetByID(uint(tracking.PasienID))
	namaPasien := ""
	if pasien != nil {
		namaPasien = pasien.Nama
	}

	return persistence.TrackingJadwalToDTO(tracking, namaObat, namaPasien), nil
}

func (u *TrackingJadwalUsecase) GetByPasienID(pasienID int) ([]dto.TrackingJadwalDTO, error) {
	trackings, err := u.trackingRepo.GetByPasienID(pasienID)
	if err != nil {
		return nil, err
	}

	pasien, _ := u.pasienRepo.GetByID(uint(pasienID))
	namaPasien := ""
	if pasien != nil {
		namaPasien = pasien.Nama
	}

	var responses []dto.TrackingJadwalDTO
	for _, tracking := range trackings {
		jadwal, _ := u.jadwalRepo.GetByID(tracking.JadwalID)
		namaObat := ""
		if jadwal != nil {
			if jadwal.KategoriObat == "Mandiri" {
				continue
			}
			namaObat = jadwal.NamaObat
		}

		dto := persistence.TrackingJadwalToDTO(&tracking, namaObat, namaPasien)
		responses = append(responses, *dto)
	}

	// ── GENERATE DYNAMIC MISSED TRACKINGS ──
	jadwals, err := u.jadwalRepo.GetByPasienID(pasienID)
	if err == nil {
		var filteredJadwals []domain.Jadwal
		for _, j := range jadwals {
			if j.KategoriObat != "Mandiri" {
				filteredJadwals = append(filteredJadwals, j)
			}
		}
		missedResponses := u.generateDynamicMissedTrackings(trackings, filteredJadwals, namaPasien)
		responses = append(responses, missedResponses...)
	}

	// Sort responses by Tanggal descending
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Tanggal > responses[j].Tanggal
	})

	return responses, nil
}

// GetByPasienIDIncludingMandiri sama seperti GetByPasienID tetapi MENYERTAKAN
// tracking obat mandiri. Digunakan oleh endpoint pasien sendiri (/api/pasien/riwayat)
// agar Flutter bisa merekonstruksi status checkbox obat mandiri setelah refresh.
func (u *TrackingJadwalUsecase) GetByPasienIDIncludingMandiri(pasienID int) ([]dto.TrackingJadwalDTO, error) {
	trackings, err := u.trackingRepo.GetByPasienID(pasienID)
	if err != nil {
		return nil, err
	}

	pasien, _ := u.pasienRepo.GetByID(uint(pasienID))
	namaPasien := ""
	if pasien != nil {
		namaPasien = pasien.Nama
	}

	var responses []dto.TrackingJadwalDTO
	for _, tracking := range trackings {
		jadwal, _ := u.jadwalRepo.GetByID(tracking.JadwalID)
		namaObat := ""
		if jadwal != nil {
			namaObat = jadwal.NamaObat
			// Obat mandiri: isi field Jadwal dengan waktu_minum slot agar
			// Flutter bisa mencocokkan slotKey = "jadwalId_waktu"
		}

		dto := persistence.TrackingJadwalToDTO(&tracking, namaObat, namaPasien)
		responses = append(responses, *dto)
	}

	// ── GENERATE DYNAMIC MISSED TRACKINGS (non-mandiri only) ──
	jadwals, err := u.jadwalRepo.GetByPasienID(pasienID)
	if err == nil {
		var filteredJadwals []domain.Jadwal
		for _, j := range jadwals {
			if j.KategoriObat != "Mandiri" {
				filteredJadwals = append(filteredJadwals, j)
			}
		}
		missedResponses := u.generateDynamicMissedTrackings(trackings, filteredJadwals, namaPasien)
		responses = append(responses, missedResponses...)
	}

	// Sort responses by Tanggal descending
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Tanggal > responses[j].Tanggal
	})

	return responses, nil
}

func (u *TrackingJadwalUsecase) Create(req *dto.CreateTrackingJadwalDTO) (*dto.TrackingJadwalDTO, error) {
	if req.JadwalID == 0 || req.PasienID == 0 {
		return nil, errors.New("jadwal_id dan pasien_id harus disediakan")
	}

	tanggal, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil {
		return nil, errors.New("format tanggal tidak valid, gunakan YYYY-MM-DD")
	}

	// ── AUTO-DETERMINE STATUS via Compliance Checker ──────────────
	// Status ditentukan server berdasarkan waktu konfirmasi vs jadwal.
	// Jika WaktuMinum kosong, gunakan time.Now() sebagai waktu konfirmasi.
	wibLoc, locErr := utils.GetWIBLocation()
	if locErr != nil {
		wibLoc = time.Local
	}
	
	confirmationTime := time.Now().In(wibLoc)
	if req.WaktuMinum != "" {
		// Gabungkan tanggal + waktu_minum untuk mendapat timestamp lengkap
		parsed, parseErr := time.ParseInLocation(
			"2006-01-02 15:04",
			req.Tanggal+" "+req.WaktuMinum,
			wibLoc,
		)
		if parseErr == nil {
			confirmationTime = parsed
		}
	}

	// Ambil jadwal untuk mendapatkan scheduled_time
	schedStatus := req.Status // fallback ke status dari request
	jadwal, jadwalErr := u.jadwalRepo.GetByID(req.JadwalID)
	if jadwalErr == nil && jadwal != nil && jadwal.WaktuMinum != "" {
		// Cari slot waktu yang sesuai jika ada banyak (dipisah koma)
		targetSlot := req.WaktuMinum
		if targetSlot == "" {
			// Jika dari request kosong, gunakan slot pertama sebagai fallback
			parts := strings.Split(jadwal.WaktuMinum, ",")
			if len(parts) > 0 {
				targetSlot = strings.TrimSpace(parts[0])
			}
		} else {
			// Validasi apakah targetSlot ada di jadwal.WaktuMinum
			found := false
			for _, part := range strings.Split(jadwal.WaktuMinum, ",") {
				if strings.TrimSpace(part) == targetSlot {
					found = true
					break
				}
			}
			if !found {
				// Jika waktu minum tidak ada di jadwal, mungkin ini jadwal yang beda, tapi fallback aja
				targetSlot = ""
			}
		}

		if targetSlot != "" {
			result, checkErr := utils.CheckComplianceFromStrings(
				req.Tanggal,
				targetSlot,
				confirmationTime,
				wibLoc,
			)
			if checkErr == nil {
				schedStatus = string(result.Status)
			}
		}
	}
	// ─────────────────────────────────────────────────────────────

	tracking := &domain.TrackingJadwal{
		JadwalID:   req.JadwalID,
		PasienID:   req.PasienID,
		Tanggal:    tanggal,
		Status:     schedStatus,
		WaktuMinum: req.WaktuMinum,
		Catatan:    req.Catatan,
		BuktiFoto:  req.BuktiFoto,
	}

	result2, err := u.trackingRepo.Create(tracking)
	if err != nil {
		return nil, err
	}

	namaObat := ""
	if jadwal != nil {
		namaObat = jadwal.NamaObat
	} else {
		jadwal2, _ := u.jadwalRepo.GetByID(result2.JadwalID)
		if jadwal2 != nil {
			namaObat = jadwal2.NamaObat
		}
	}

	pasien, _ := u.pasienRepo.GetByID(uint(result2.PasienID))
	namaPasien := ""
	if pasien != nil {
		namaPasien = pasien.Nama
	}

	return persistence.TrackingJadwalToDTO(result2, namaObat, namaPasien), nil
}

func (u *TrackingJadwalUsecase) Update(id int, req *dto.UpdateTrackingJadwalDTO) (*dto.TrackingJadwalDTO, error) {
	tracking, err := u.trackingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Status != "" {
		tracking.Status = req.Status
	}
	if req.WaktuMinum != "" {
		tracking.WaktuMinum = req.WaktuMinum
	}
	if req.Catatan != "" {
		tracking.Catatan = req.Catatan
	}
	if req.BuktiFoto != "" {
		tracking.BuktiFoto = req.BuktiFoto
	}

	result, err := u.trackingRepo.Update(id, tracking)
	if err != nil {
		return nil, err
	}

	jadwal, _ := u.jadwalRepo.GetByID(result.JadwalID)
	namaObat := ""
	if jadwal != nil {
		namaObat = jadwal.NamaObat
	}

	pasien, _ := u.pasienRepo.GetByID(uint(result.PasienID))
	namaPasien := ""
	if pasien != nil {
		namaPasien = pasien.Nama
	}

	return persistence.TrackingJadwalToDTO(result, namaObat, namaPasien), nil
}

func (u *TrackingJadwalUsecase) Delete(id int) error {
	return u.trackingRepo.Delete(id)
}

// GetStatistics untuk dashboard compliance
func (u *TrackingJadwalUsecase) GetStatistics(pasienID int) (map[string]interface{}, error) {
	var trackings []domain.TrackingJadwal
	var err error

	if pasienID > 0 {
		trackings, err = u.trackingRepo.GetByPasienID(pasienID)
	} else {
		trackings, err = u.trackingRepo.GetAll()
	}

	if err != nil {
		return nil, err
	}

	// Filter out trackings related to "Mandiri" jadwals
	var filteredTrackings []domain.TrackingJadwal
	for _, t := range trackings {
		jadwal, _ := u.jadwalRepo.GetByID(t.JadwalID)
		if jadwal != nil && jadwal.KategoriObat == "Mandiri" {
			continue
		}
		filteredTrackings = append(filteredTrackings, t)
	}
	trackings = filteredTrackings

	diminum := 0
	terlambat := 0
	terlewat := 0

	for _, t := range trackings {
		switch t.Status {
		case "Diminum":
			diminum++
		case "Terlambat":
			terlambat++
		case "Terlewat":
			terlewat++
		}
	}

	// Fetch schedules to generate dynamic "Terlewat" count
	var jadwals []domain.Jadwal
	if pasienID > 0 {
		jadwals, _ = u.jadwalRepo.GetByPasienID(pasienID)
		// Filter out mandiri jadwals
		var filteredJadwals []domain.Jadwal
		for _, j := range jadwals {
			if j.KategoriObat != "Mandiri" {
				filteredJadwals = append(filteredJadwals, j)
			}
		}
		jadwals = filteredJadwals

		pasien, _ := u.pasienRepo.GetByID(uint(pasienID))
		namaPasien := ""
		if pasien != nil {
			namaPasien = pasien.Nama
		}
		missed := u.generateDynamicMissedTrackings(trackings, jadwals, namaPasien)
		terlewat += len(missed)
	} else {
		jadwals, _ = u.jadwalRepo.GetAll()
		// Filter out mandiri jadwals
		var filteredJadwals []domain.Jadwal
		for _, j := range jadwals {
			if j.KategoriObat != "Mandiri" {
				filteredJadwals = append(filteredJadwals, j)
			}
		}
		jadwals = filteredJadwals

		trackingsByPasien := make(map[int][]domain.TrackingJadwal)
		for _, t := range trackings {
			trackingsByPasien[t.PasienID] = append(trackingsByPasien[t.PasienID], t)
		}

		jadwalsByPasien := make(map[int][]domain.Jadwal)
		for _, j := range jadwals {
			jadwalsByPasien[j.PasienID] = append(jadwalsByPasien[j.PasienID], j)
		}

		for pID, pJadwals := range jadwalsByPasien {
			pasien, _ := u.pasienRepo.GetByID(uint(pID))
			namaPasien := ""
			if pasien != nil {
				namaPasien = pasien.Nama
			}
			pTrackings := trackingsByPasien[pID]
			missed := u.generateDynamicMissedTrackings(pTrackings, pJadwals, namaPasien)
			terlewat += len(missed)
		}
	}

	total := diminum + terlambat + terlewat

	stats := map[string]interface{}{
		"diminum":         diminum,
		"terlambat":       terlambat,
		"terlewat":        terlewat,
		"total":           total,
		"compliance_rate": 0.0,
	}

	if total > 0 {
		complianceRate := float64(diminum) / float64(total) * 100
		stats["compliance_rate"] = complianceRate
	}

	return stats, nil
}

func (u *TrackingJadwalUsecase) generateDynamicMissedTrackings(trackings []domain.TrackingJadwal, jadwals []domain.Jadwal, namaPasien string) []dto.TrackingJadwalDTO {
	loc := time.FixedZone("WIB", 7*60*60)
	now := time.Now().In(loc)

	existingLogs := make(map[string]bool) // key: "jadwalID_YYYY-MM-DD"
	var responses []dto.TrackingJadwalDTO

	for _, tracking := range trackings {
		tanggalStr := tracking.Tanggal.In(loc).Format("2006-01-02")
		key := fmt.Sprintf("%d_%s", tracking.JadwalID, tanggalStr)
		existingLogs[key] = true
	}

	for _, j := range jadwals {
		if j.WaktuMinum == "" {
			continue
		}

		startDate, err := time.ParseInLocation("2006-01-02", j.TanggalMulai, loc)
		if err != nil {
			continue
		}

		endDate := now
		if j.TanggalSelesai != "" {
			parsedEnd, err := time.ParseInLocation("2006-01-02", j.TanggalSelesai, loc)
			if err == nil {
				endDate = parsedEnd
			}
		}

		limitDate := endDate
		if limitDate.After(now) {
			limitDate = now
		}

		// Normalize limitDate and currentDate to midnight
		limitDate = time.Date(limitDate.Year(), limitDate.Month(), limitDate.Day(), 0, 0, 0, 0, loc)
		currentDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, loc)

		waktuList := []string{j.WaktuMinum}
		if strings.Contains(j.WaktuMinum, ",") {
			parts := strings.Split(j.WaktuMinum, ",")
			waktuList = []string{}
			for _, part := range parts {
				waktuList = append(waktuList, strings.TrimSpace(part))
			}
		}

		for !currentDate.After(limitDate) {
			dateStr := currentDate.Format("2006-01-02")

			for _, wm := range waktuList {
				if wm == "" {
					continue
				}

				key := fmt.Sprintf("%d_%s", j.JadwalID, dateStr)
				if existingLogs[key] {
					continue
				}

				schedTimeFull, err := time.ParseInLocation("2006-01-02 15:04", dateStr+" "+wm, loc)
				if err != nil {
					continue
				}

				// If it's in the past by more than 75 minutes, generate dynamic "Terlewat" log
				if now.Sub(schedTimeFull).Minutes() > 75 {
					responses = append(responses, dto.TrackingJadwalDTO{
						ID:         0,
						JadwalID:   j.JadwalID,
						PasienID:   j.PasienID,
						Tanggal:    dateStr,
						Jadwal:     wm,
						Status:     "Terlewat",
						WaktuMinum: "",
						Catatan:    "Terlewat minum obat (sistem otomatis)",
						NamaObat:   j.NamaObat,
						NamaPasien: namaPasien,
					})
				}
			}

			currentDate = currentDate.AddDate(0, 0, 1)
		}
	}

	return responses
}

// GetObatStreak menghitung streak kepatuhan meminum obat pasien
func (u *TrackingJadwalUsecase) GetObatStreak(pasienID int) (int, error) {
	trackings, err := u.trackingRepo.GetByPasienID(pasienID)
	if err != nil {
		return 0, err
	}

	// Buat map tanggal -> status list
	trackingByDate := make(map[string][]string)
	wib := time.FixedZone("WIB", 7*60*60)
	for _, t := range trackings {
		tgl := t.Tanggal.In(wib).Format("2006-01-02")
		trackingByDate[tgl] = append(trackingByDate[tgl], t.Status)
	}

	streak := 0
	dateToCheck := time.Now().In(wib)

	for {
		tgl := dateToCheck.Format("2006-01-02")
		statuses, ok := trackingByDate[tgl]

		if !ok {
			// Jika tidak ada data hari ini, mungkin belum jadwalnya.
			// Jangan putus streak untuk hari ini. Tapi jika hari kemarin tidak ada, putus.
			if tgl == time.Now().In(wib).Format("2006-01-02") {
				dateToCheck = dateToCheck.AddDate(0, 0, -1)
				continue
			}
			break
		}

		hasTerlewat := false
		hasDiminum := false
		for _, s := range statuses {
			if s == "Terlewat" {
				hasTerlewat = true
			}
			if s == "Diminum" || s == "Terlambat" {
				hasDiminum = true
			}
		}

		if hasTerlewat {
			break
		}

		if hasDiminum {
			streak++
			dateToCheck = dateToCheck.AddDate(0, 0, -1)
		} else {
			break
		}
	}

	return streak, nil
}

