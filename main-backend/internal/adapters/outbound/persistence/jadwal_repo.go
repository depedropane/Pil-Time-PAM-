package persistence

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"

	"gorm.io/gorm"
)

var _ outbound.JadwalRepository = (*jadwalRepoImpl)(nil)

type jadwalRepoImpl struct {
	db *gorm.DB
}

func NewJadwalRepo(db *gorm.DB) outbound.JadwalRepository {
	return &jadwalRepoImpl{db}
}

func (r *jadwalRepoImpl) GetAll() ([]domain.Jadwal, error) {
	var jadwals []domain.Jadwal
	err := r.db.
		Table("jadwal j").
		Select("j.*, p.nama as pasien_nama").
		Joins("LEFT JOIN pasien p ON j.pasien_id = p.pasien_id").
		Scan(&jadwals).Error
	return jadwals, err
}

func (r *jadwalRepoImpl) GetByID(id int) (*domain.Jadwal, error) {
	var jadwal domain.Jadwal
	err := r.db.
		Table("jadwal j").
		Select("j.*, p.nama as pasien_nama").
		Joins("LEFT JOIN pasien p ON j.pasien_id = p.pasien_id").
		Where("j.jadwal_id = ?", id).
		Scan(&jadwal).Error
	if err != nil {
		return nil, err
	}
	return &jadwal, nil
}

func (r *jadwalRepoImpl) GetByPasienID(pasienID int) ([]domain.Jadwal, error) {
	var jadwals []domain.Jadwal
	err := r.db.
		Table("jadwal j").
		Select("j.*, p.nama as pasien_nama").
		Joins("LEFT JOIN pasien p ON j.pasien_id = p.pasien_id").
		Where("j.pasien_id = ?", pasienID).
		Scan(&jadwals).Error
	return jadwals, err
}

func (r *jadwalRepoImpl) Create(j *domain.Jadwal) (*domain.Jadwal, error) {
	err := r.db.Create(j).Error
	return j, err
}

func (r *jadwalRepoImpl) Update(id int, j *domain.Jadwal) (*domain.Jadwal, error) {
	updateMap := map[string]interface{}{
		"nama_obat":            j.NamaObat,
		"jumlah_dosis":         j.JumlahDosis,
		"satuan":               j.Satuan,
		"kategori_obat":        j.KategoriObat,
		"takaran_obat":         j.TakaranObat,
		"frekuensi_per_hari":   j.FrekuensiPerHari,
		"waktu_minum":          j.WaktuMinum,
		"aturan_konsumsi":      j.AturanKonsumsi,
		"catatan":              j.Catatan,
		"tipe_durasi":          j.TipeDurasi,
		"jumlah_hari":          j.JumlahHari,
		"tanggal_mulai":        j.TanggalMulai,
		"tanggal_selesai":      j.TanggalSelesai,
		"status":               j.Status,
	}
	err := r.db.Model(&domain.Jadwal{}).Where("jadwal_id = ?", id).Updates(updateMap).Error
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *jadwalRepoImpl) Delete(id int) error {
	return r.db.Delete(&domain.Jadwal{}, id).Error
}

// JadwalToResponseDTO converts domain Jadwal to JadwalResponseDTO.
// PasienNama diambil dari field domain.Jadwal.PasienNama (hasil JOIN).
func JadwalToResponseDTO(jadwal *domain.Jadwal) *dto.JadwalResponseDTO {
	return &dto.JadwalResponseDTO{
		ID:                 jadwal.JadwalID,
		PasienID:           jadwal.PasienID,
		PasienNama:         jadwal.PasienNama,
		NamaObat:           jadwal.NamaObat,
		JumlahDosis:        jadwal.JumlahDosis,
		Satuan:             jadwal.Satuan,
		KategoriObat:       jadwal.KategoriObat,
		TakaranObat:        jadwal.TakaranObat,
		FrekuensiPerHari:   jadwal.FrekuensiPerHari,
		WaktuMinum:         jadwal.WaktuMinum,
		AturanKonsumsi:     jadwal.AturanKonsumsi,
		Catatan:            jadwal.Catatan,
		TipeDurasi:         jadwal.TipeDurasi,
		JumlahHari:         jadwal.JumlahHari,
		TanggalMulai:       jadwal.TanggalMulai,
		TanggalSelesai:     jadwal.TanggalSelesai,
		Status:             jadwal.Status,
		CreatedAt:          jadwal.CreatedAt,
		UpdatedAt:          jadwal.UpdatedAt,
	}
}

func (r *jadwalRepoImpl) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Jadwal{}).Count(&count).Error
	return count, err
}
