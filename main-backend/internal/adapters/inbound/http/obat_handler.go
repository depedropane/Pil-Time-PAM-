package http

import (
    "backend/internal/dto"
    "backend/internal/usecase"
    "fmt" // <--- TAMBAHKAN INI
    "net/http"
    "os"
    "path/filepath"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type ObatHandler struct {
	usecase *usecase.ObatUsecase
}

func NewObatHandler(u *usecase.ObatUsecase) *ObatHandler {
	return &ObatHandler{usecase: u}
}

// GetAll mendapatkan semua obat (admin)
func (h *ObatHandler) GetAll(c *gin.Context) {
	obats, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "FETCH_ERROR",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": obats})
}

// GetAllForPasien mendapatkan obat mandiri milik pasien yang sedang login (dari JWT)
func (h *ObatHandler) GetAllForPasien(c *gin.Context) {
	pasienID := c.GetInt("pasien_id")
	if pasienID <= 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_PARAMETER",
			Message: "pasien_id tidak ditemukan di token",
		})
		return
	}
	obats, err := h.usecase.GetAllByPasien(pasienID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "FETCH_ERROR",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": obats})
}

// Create menangani FormData (Text + File)
func (h *ObatHandler) Create(c *gin.Context) {
    var req dto.CreateObatDTO

    // 1. Bind data teks terlebih dahulu
    if err := c.ShouldBind(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.ErrorResponse{
            Error: "VALIDATION_ERROR", Message: err.Error(),
        })
        return
    }

    // 2. Ambil file gambar (Key: gambar_file sesuai di Vue kamu)
    file, err := c.FormFile("gambar_file")
    if err == nil {
        uploadDir := "public/uploads"
        os.MkdirAll(uploadDir, os.ModePerm)

        filename := uuid.New().String() + filepath.Ext(file.Filename)
        filePath := filepath.Join(uploadDir, filename)

        if err := c.SaveUploadedFile(file, filePath); err == nil {
            // SETELAH Save sukses, baru masukkan path ke req.Gambar
            req.Gambar = "/uploads/" + filename
        }
    } else {
        fmt.Println("Info: Tidak ada file gambar yang diupload:", err)
    }

    // 3. Simpan ke Database via Usecase
    resp, err := h.usecase.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "CREATE_ERROR", Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

// Update data obat
func (h *ObatHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.UpdateObatDTO

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "VALIDATION_ERROR", Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("gambar_file")
	if err == nil {
		uploadDir := "public/uploads"
		os.MkdirAll(uploadDir, os.ModePerm)
		filename := uuid.New().String() + filepath.Ext(file.Filename)
		filePath := filepath.Join(uploadDir, filename)
		c.SaveUploadedFile(file, filePath)
		req.Gambar = "/uploads/" + filename
	}

	resp, err := h.usecase.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "UPDATE_ERROR", Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetByID mendapatkan satu obat
func (h *ObatHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	obat, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "NOT_FOUND", Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": obat})
}

// Delete obat
func (h *ObatHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	obat, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Obat tidak ditemukan",
		})
		return
	}

	// Hanya boleh menghapus Obat Mandiri, Obat Master dilarang dihapus
	if !obat.IsMandiri {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{
			Error:   "FORBIDDEN",
			Message: "Obat adalah data master dan tidak dapat dihapus",
		})
		return
	}

	err = h.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "DELETE_ERROR",
			Message: "Gagal menghapus obat",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Obat mandiri berhasil dihapus"})
}

// CreateMandiri menangani pembuatan obat mandiri (independent medication)
func (h *ObatHandler) CreateMandiri(c *gin.Context) {
	var req dto.CreateObatMandiriDTO

	// Bind data form/JSON
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	// Ambil file gambar jika ada
	file, err := c.FormFile("gambar_file")
	if err == nil {
		uploadDir := "public/uploads"
		os.MkdirAll(uploadDir, os.ModePerm)

		filename := uuid.New().String() + filepath.Ext(file.Filename)
		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, filePath); err == nil {
			req.Gambar = "/uploads/" + filename
		}
	}

	// Simpan ke Database via Usecase
	resp, err := h.usecase.CreateMandiri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "CREATE_ERROR",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": resp})
}