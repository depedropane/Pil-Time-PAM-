package persistence

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"
	"time"

	"gorm.io/gorm"
)

var _ outbound.TrackingJadwalRepository = (*trackingJadwalRepoImpl)(nil)
var _ outbound.RiwayatObatRepository = (*riwayatObatRepoImpl)(nil)

// ==================== TrackingJadwal Repository ====================

type trackingJadwalRepoImpl struct {
	db *gorm.DB
}

func NewTrackingJadwalRepo(db *gorm.DB) outbound.TrackingJadwalRepository {
	return &trackingJadwalRepoImpl{db}
}

func (r *trackingJadwalRepoImpl) GetAll() ([]domain.TrackingJadwal, error) {
	var trackings []domain.TrackingJadwal
	err := r.db.Order("tanggal DESC").Find(&trackings).Error
	return trackings, err
}

func (r *trackingJadwalRepoImpl) GetByID(id int) (*domain.TrackingJadwal, error) {
	var tracking domain.TrackingJadwal
	err := r.db.First(&tracking, id).Error
	return &tracking, err
}

func (r *trackingJadwalRepoImpl) GetByPasienID(pasienID int) ([]domain.TrackingJadwal, error) {
	var trackings []domain.TrackingJadwal
	err := r.db.Where("pasien_id = ?", pasienID).Order("tanggal DESC").Find(&trackings).Error
	return trackings, err
}

func (r *trackingJadwalRepoImpl) GetByJadwalID(jadwalID int) ([]domain.TrackingJadwal, error) {
	var trackings []domain.TrackingJadwal
	err := r.db.Where("jadwal_id = ?", jadwalID).Order("tanggal DESC").Find(&trackings).Error
	return trackings, err
}

func (r *trackingJadwalRepoImpl) GetByTanggal(tanggal string) ([]domain.TrackingJadwal, error) {
	var trackings []domain.TrackingJadwal
	err := r.db.Where("DATE(tanggal) = ?", tanggal).Order("tanggal DESC").Find(&trackings).Error
	return trackings, err
}

func (r *trackingJadwalRepoImpl) Create(tracking *domain.TrackingJadwal) (*domain.TrackingJadwal, error) {
	err := r.db.Create(tracking).Error
	return tracking, err
}

func (r *trackingJadwalRepoImpl) Update(id int, tracking *domain.TrackingJadwal) (*domain.TrackingJadwal, error) {
	err := r.db.Model(&domain.TrackingJadwal{}).Where("tracking_jadwal_id = ?", id).Updates(tracking).Error
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *trackingJadwalRepoImpl) Delete(id int) error {
	return r.db.Delete(&domain.TrackingJadwal{}, id).Error
}

// Helper to convert domain to DTO
func TrackingJadwalToDTO(tracking *domain.TrackingJadwal, namaObat, namaPasien string) *dto.TrackingJadwalDTO {
	wib := time.FixedZone("WIB", 7*60*60)
	tanggalWIB := tracking.Tanggal.In(wib)

	jadwalStr := ""
	if tracking.Tanggal != (time.Time{}) {
		jadwalStr = tanggalWIB.Format("15:04")
	}

	return &dto.TrackingJadwalDTO{
		ID:         tracking.TrackingJadwalID,
		JadwalID:   tracking.JadwalID,
		PasienID:   tracking.PasienID,
		NamaPasien: namaPasien,
		NamaObat:   namaObat,
		Tanggal:    tanggalWIB.Format("2006-01-02"),
		Jadwal:     jadwalStr,
		WaktuMinum: tracking.WaktuMinum,
		Status:     tracking.Status,
		Catatan:    tracking.Catatan,
		BuktiFoto:  tracking.BuktiFoto,
		CreatedAt:  tracking.CreatedAt,
	}
}


// ==================== RiwayatObat Repository ====================

type riwayatObatRepoImpl struct {
	db *gorm.DB
}

func NewRiwayatObatRepo(db *gorm.DB) outbound.RiwayatObatRepository {
	return &riwayatObatRepoImpl{db}
}

func (r *riwayatObatRepoImpl) GetAll() ([]domain.RiwayatObat, error) {
	var riwayats []domain.RiwayatObat
	err := r.db.Order("tanggal DESC").Find(&riwayats).Error
	return riwayats, err
}

func (r *riwayatObatRepoImpl) GetByID(id int) (*domain.RiwayatObat, error) {
	var riwayat domain.RiwayatObat
	err := r.db.First(&riwayat, id).Error
	return &riwayat, err
}

func (r *riwayatObatRepoImpl) GetByPasienID(pasienID int) ([]domain.RiwayatObat, error) {
	var riwayats []domain.RiwayatObat
	err := r.db.Where("pasien_id = ?", pasienID).Order("tanggal DESC").Find(&riwayats).Error
	return riwayats, err
}

func (r *riwayatObatRepoImpl) Create(riwayat *domain.RiwayatObat) (*domain.RiwayatObat, error) {
	err := r.db.Create(riwayat).Error
	return riwayat, err
}

func (r *riwayatObatRepoImpl) Update(id int, riwayat *domain.RiwayatObat) (*domain.RiwayatObat, error) {
	err := r.db.Model(&domain.RiwayatObat{}).Where("riwayat_obat_id = ?", id).Updates(riwayat).Error
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *riwayatObatRepoImpl) Delete(id int) error {
	return r.db.Delete(&domain.RiwayatObat{}, id).Error
}

// Helper to convert domain to DTO
func RiwayatObatToDTO(riwayat *domain.RiwayatObat, namaPasien string) *dto.RiwayatObatDTO {
	return &dto.RiwayatObatDTO{
		ID:         riwayat.RiwayatObatID,
		PasienID:   riwayat.PasienID,
		NamaPasien: namaPasien,
		Tanggal:    riwayat.Tanggal.Format("2006-01-02"),
		Status:     riwayat.Status,
		Catatan:    riwayat.Catatan,
		CreatedAt:  riwayat.CreatedAt,
	}
}
