package dto

type ObatDosisDTO struct {
	ObatID int    `json:"obat_id"`
	Dosis  string `json:"dosis"`
}

type CreateResepWithJadwalDTO struct {
    PasienID int `json:"pasien_id"`
    NakesID  int `json:"nakes_id"`

    ObatList []ObatDosisDTO `json:"obat_list"`

    Catatan string `json:"catatan"`

    TanggalMulai   string `json:"tanggal_mulai"`
    TanggalSelesai string `json:"tanggal_selesai"`

    FrekuensiPerHari int      `json:"frekuensi_per_hari"`
    AturanKonsumsi   string   `json:"aturan_konsumsi"`
    JamMinum         []string `json:"jam_minum"`
    
    TipeDurasi       string   `json:"tipe_durasi"`
    JumlahHari       int      `json:"jumlah_hari"`
}