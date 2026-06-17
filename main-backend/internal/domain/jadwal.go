package domain

import "time"

type Jadwal struct {
	JadwalID           int       `gorm:"primaryKey;column:jadwal_id"`
	PasienID           int       `gorm:"column:pasien_id;not null;index"`
	PasienNama         string    `gorm:"->;"`                             // Diisi dari JOIN, tidak disimpan ke tabel jadwal
	NamaObat           string    `gorm:"column:nama_obat;not null"`
	JumlahDosis        int       `gorm:"column:jumlah_dosis;not null"`
	Satuan             string    `gorm:"column:satuan;not null"`
	KategoriObat       string    `gorm:"column:kategori_obat;not null"`
	TakaranObat        string    `gorm:"column:takaran_obat;not null"`
	FrekuensiPerHari   string    `gorm:"column:frekuensi_per_hari;not null"`
	WaktuMinum         string    `gorm:"column:waktu_minum;not null"`
	AturanKonsumsi     string    `gorm:"column:aturan_konsumsi;not null"`
	Catatan            string    `gorm:"column:catatan;type:text"`
	TipeDurasi         string    `gorm:"column:tipe_durasi;not null"` // 'hari' or 'rutin'
	JumlahHari         int       `gorm:"column:jumlah_hari"`
	TanggalMulai       string    `gorm:"column:tanggal_mulai"`
	TanggalSelesai     string    `gorm:"column:tanggal_selesai"`
	Status             string    `gorm:"column:status;not null"` // 'aktif', 'rutin', etc
	CreatedAt          time.Time `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
}

func (Jadwal) TableName() string {
	return "jadwal"
}
