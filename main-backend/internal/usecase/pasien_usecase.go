package usecase

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/ports/outbound"
	"backend/pkg/email"
	"backend/pkg/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PasienUsecase struct {
	repo         outbound.PasienRepository
	jadwalRepo   outbound.JadwalRepository
	emailService *email.EmailService
}

func NewPasienUsecase(r outbound.PasienRepository, j outbound.JadwalRepository) *PasienUsecase {
	return &PasienUsecase{
		repo:         r,
		jadwalRepo:   j,
		emailService: email.NewEmailService(),
	}
}

// Register melakukan registrasi pasien baru
func (u *PasienUsecase) Register(req *dto.RegisterPasienRequest) (*dto.RegisterPasienResponse, error) {
	// Cek email sudah ada atau belum
	existingEmail, _ := u.repo.GetByEmail(req.Email)
	if existingEmail != nil && existingEmail.Email != "" {
		return nil, errors.New("email sudah terdaftar")
	}

	// Cek NIK sudah ada atau belum
	existingNIK, _ := u.repo.GetByNIK(req.NIK)
	if existingNIK != nil && existingNIK.NIK != "" {
		return nil, errors.New("NIK sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	// Parse tanggal lahir
	tanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		return nil, errors.New("format tanggal lahir salah (gunakan YYYY-MM-DD)")
	}

	// Buat pasien baru
	pasien := &domain.Pasien{
		Nama:         req.NamaLengkap,
		Email:        req.Email,
		Password:     hashedPassword,
		NIK:          req.NIK,
		TanggalLahir: tanggalLahir,
		NoTelepon:           req.Telepon,
		NoTeleponPendamping: req.NoTeleponPendamping,
		JenisKelamin:        req.JenisKelamin,
		Alamat:              req.Alamat,
	}

	// Simpan ke database
	if err := u.repo.Create(pasien); err != nil {
		return nil, errors.New("gagal menyimpan data pasien")
	}

	// Return response
	return &dto.RegisterPasienResponse{
		PasienID:    pasien.PasienID,
		Email:       pasien.Email,
		NamaLengkap: pasien.Nama,
		NIK:         pasien.NIK,
		Message:     "Registrasi berhasil",
	}, nil
}

func (u *PasienUsecase) GetAll() ([]domain.Pasien, error) {
	return u.repo.GetAll()
}

// GetJadwalByPasien mengambil daftar jadwal untuk pasien tertentu
func (u *PasienUsecase) GetJadwalByPasien(pasienID int) ([]domain.Jadwal, error) {
	return u.jadwalRepo.GetByPasienID(pasienID)
}

// Login melakukan login pasien
func (u *PasienUsecase) Login(req *dto.LoginPasienRequest) (*dto.LoginPasienResponse, error) {
	// Cari pasien berdasarkan email
	pasien, err := u.repo.GetByEmail(req.Email)
	if err != nil || pasien == nil || pasien.Email == "" {
		return nil, errors.New("email atau password salah")
	}

	// Verify password
	if err := utils.VerifyPassword(req.Password, pasien.Password); err != nil {
		return nil, errors.New("email atau password salah")
	}

	// Return response
	return &dto.LoginPasienResponse{
		PasienID:    pasien.PasienID,
		Email:       pasien.Email,
		NamaLengkap: pasien.Nama,
		Message:     "Login berhasil",
	}, nil
}

// ForgotPassword mengirim kode reset password ke email
func (u *PasienUsecase) ForgotPassword(req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, error) {
	// Cek apakah email terdaftar
	pasien, err := u.repo.GetByEmail(req.Email)
	if err != nil || pasien == nil || pasien.Email == "" {
		return nil, errors.New("email tidak terdaftar")
	}

	// Generate 6-digit random code
	code := generateRandomCode()

	// Set expiry time (15 menit)
	expiryTime := time.Now().Add(15 * time.Minute)

	// Save reset code ke database
	if err := u.repo.UpdateResetCode(req.Email, code, expiryTime); err != nil {
		return nil, errors.New("gagal menyimpan kode reset")
	}

	// Kirim email dengan kode reset
	if err := u.emailService.SendResetCode(req.Email, code); err != nil {
		fmt.Printf("[RESEND ERROR] Gagal mengirim email ke %s: %v\n", req.Email, err)
		return nil, errors.New("gagal mengirim email reset password")
	}

	return &dto.ForgotPasswordResponse{
		Message: "Kode reset password telah dikirim ke email Anda",
	}, nil
}

// VerifyResetCode memverifikasi kode reset password
func (u *PasienUsecase) VerifyResetCode(req *dto.VerifyResetCodeRequest) (*dto.VerifyResetCodeResponse, error) {
	// Cek apakah email terdaftar
	pasien, err := u.repo.GetByEmail(req.Email)
	if err != nil || pasien == nil || pasien.Email == "" {
		return nil, errors.New("email tidak terdaftar")
	}

	// Cek apakah reset code ada
	if pasien.ResetCode == nil || *pasien.ResetCode == "" {
		return nil, errors.New("tidak ada permintaan reset password untuk email ini")
	}

	// Cek apakah kode sudah expired
	if pasien.ResetCodeExpiry == nil || time.Now().After(*pasien.ResetCodeExpiry) {
		return nil, errors.New("kode reset password sudah kadaluarsa")
	}

	// Cek apakah kode sesuai
	if *pasien.ResetCode != req.Code {
		return nil, errors.New("kode reset password tidak valid")
	}

	return &dto.VerifyResetCodeResponse{
		Message: "Kode reset password valid",
	}, nil
}

// ResetPassword mengubah password pasien
func (u *PasienUsecase) ResetPassword(req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error) {
	// Cek apakah email terdaftar
	pasien, err := u.repo.GetByEmail(req.Email)
	if err != nil || pasien == nil || pasien.Email == "" {
		return nil, errors.New("email tidak terdaftar")
	}

	// Verifikasi kode terlebih dahulu
	verifyReq := &dto.VerifyResetCodeRequest{
		Email: req.Email,
		Code:  req.Code,
	}
	if _, err := u.VerifyResetCode(verifyReq); err != nil {
		return nil, err
	}

	// Hash password baru
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	// Update password dan clear reset code
	if err := u.repo.UpdatePassword(req.Email, hashedPassword); err != nil {
		return nil, errors.New("gagal mengubah password")
	}

	return &dto.ResetPasswordResponse{
		Message: "Password berhasil diubah",
	}, nil
}

// generateRandomCode generate 6-digit random code
func generateRandomCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // Generate number between 100000-999999
	return fmt.Sprintf("%06d", code)
}

// GetPasienDashboard mengambil data dashboard untuk pasien dengan jadwal hari ini
func (u *PasienUsecase) GetPasienDashboard(pasienID int) (*dto.PasienDashboardResponse, error) {
	// Get pasien data
	pasien, err := u.repo.GetByID(uint(pasienID))
	if err != nil || pasien == nil {
		return nil, errors.New("pasien tidak ditemukan")
	}

	// Get all jadwal for pasien
	allJadwals, err := u.jadwalRepo.GetByPasienID(int(pasienID))
	if err != nil {
		return nil, errors.New("gagal mengambil jadwal pasien")
	}

	// Convert domain jadwal to DTO
	jadwalDTOs := []dto.JadwalResponseDTO{}
	todayJadwalDTOs := []dto.JadwalResponseDTO{}
	today := time.Now().Format("2006-01-02")

	for _, jadwal := range allJadwals {
		jadwalDTO := dto.JadwalResponseDTO{
			ID:                 jadwal.JadwalID,
			PasienID:           jadwal.PasienID,
			PasienNama:         pasien.Nama,
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

		jadwalDTOs = append(jadwalDTOs, jadwalDTO)

		// Check if jadwal is for today and still active
		if jadwal.Status == "aktif" || jadwal.Status == "active" {
			jadwalMulai := jadwal.TanggalMulai
			jadwalSelesai := jadwal.TanggalSelesai

			// Handle jadwal berdasarkan tipe_durasi
			switch jadwal.TipeDurasi {
			case "rutin":
				// Jadwal rutin berlaku setiap hari jika sudah mulai
				if jadwalMulai <= today {
					todayJadwalDTOs = append(todayJadwalDTOs, jadwalDTO)
				}
			case "hari":
				// Jadwal terbatas berlaku hingga tanggal_selesai (jika ada)
				// Jika tanggal_selesai kosong, anggap jadwal berlaku indefinite
				if jadwalMulai <= today {
					if jadwalSelesai == "" {
						// Tanpa batas akhir - anggap berlaku selamanya
						todayJadwalDTOs = append(todayJadwalDTOs, jadwalDTO)
					} else if today <= jadwalSelesai {
						// Ada batas akhir - cek apakah masih dalam rentang
						todayJadwalDTOs = append(todayJadwalDTOs, jadwalDTO)
					}
				}
			}
		}
	}

	// Return dashboard response
	return &dto.PasienDashboardResponse{
		PasienID:            pasien.PasienID,
		Nama:                pasien.Nama,
		Email:               pasien.Email,
		NoTelepon:           pasien.NoTelepon,
		NoTeleponPendamping: pasien.NoTeleponPendamping,
		JenisKelamin:        pasien.JenisKelamin,
		TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
		Alamat:       pasien.Alamat,
		TodayJadwals: todayJadwalDTOs,
		AllJadwals:   jadwalDTOs,
	}, nil
}

// GetPasienJadwal mengambil semua jadwal untuk pasien tertentu
func (u *PasienUsecase) GetPasienJadwal(pasienID int) (*dto.PasienJadwalResponse, error) {
	// Get pasien data
	pasien, err := u.repo.GetByID(uint(pasienID))
	if err != nil || pasien == nil {
		return nil, errors.New("pasien tidak ditemukan")
	}

	// Get all jadwal for pasien
	allJadwals, err := u.jadwalRepo.GetByPasienID(pasienID)
	if err != nil {
		return nil, errors.New("gagal mengambil jadwal pasien")
	}

	// Convert domain jadwal to DTO
	jadwalDTOs := []dto.JadwalResponseDTO{}
	for _, jadwal := range allJadwals {
		jadwalDTO := dto.JadwalResponseDTO{
			ID:                 jadwal.JadwalID,
			PasienID:           jadwal.PasienID,
			PasienNama:         pasien.Nama,
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
		jadwalDTOs = append(jadwalDTOs, jadwalDTO)
	}

	return &dto.PasienJadwalResponse{
		PasienID: pasien.PasienID,
		Nama:     pasien.Nama,
		Jadwals:  jadwalDTOs,
	}, nil
}

// GetByID mengambil data pasien berdasarkan ID
func (u *PasienUsecase) GetByID(pasienID int) (*dto.PasienResponseDTO, error) {
	pasien, err := u.repo.GetByID(uint(pasienID))
	if err != nil || pasien == nil {
		return nil, errors.New("pasien tidak ditemukan")
	}

	return &dto.PasienResponseDTO{
		PasienID:     pasien.PasienID,
		Nama:         pasien.Nama,
		Email:        pasien.Email,
		NIK:          pasien.NIK,
		TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
		TempatLahir:  pasien.TempatLahir,
		Alamat:       pasien.Alamat,
		JenisKelamin:        pasien.JenisKelamin,
		NoTelepon:           pasien.NoTelepon,
		NoTeleponPendamping: pasien.NoTeleponPendamping,
	}, nil
}

// UpdateProfile memperbarui profil pasien
func (u *PasienUsecase) UpdateProfile(pasienID int, req *dto.UpdatePasienRequest) error {
	existing, err := u.repo.GetByID(uint(pasienID))
	if err != nil || existing == nil {
		return errors.New("pasien tidak ditemukan")
	}

	if req.Nama != "" {
		existing.Nama = req.Nama
	}
	if req.Email != "" {
		existing.Email = req.Email
	}
	if req.Alamat != "" {
		existing.Alamat = req.Alamat
	}
	if req.NoTelepon != "" {
		existing.NoTelepon = req.NoTelepon
	}
	if req.NoTeleponPendamping != "" {
		existing.NoTeleponPendamping = req.NoTeleponPendamping
	}
	if req.JenisKelamin != "" {
		existing.JenisKelamin = req.JenisKelamin
	}
	if req.TempatLahir != "" {
		existing.TempatLahir = req.TempatLahir
	}
	if req.TanggalLahir != "" {
		parsedDate, err := time.Parse("2006-01-02", req.TanggalLahir)
		if err == nil {
			existing.TanggalLahir = parsedDate
		}
	}

	return u.repo.UpdateProfile(uint(pasienID), existing)
}
