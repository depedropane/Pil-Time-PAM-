package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"backend/internal/dto"

	"github.com/gin-gonic/gin"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// UploadImage menangani HTTP request untuk upload image
func (h *FileHandler) UploadImage(c *gin.Context) {
	// Get file dari request
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_FILE",
			Message: "Tidak dapat membaca file",
		})
		return
	}

	// Validasi ukuran file (max 5MB)
	maxSize := int64(5 * 1024 * 1024) // 5MB
	if file.Size > maxSize {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "FILE_TOO_LARGE",
			Message: "Ukuran file maksimal 5MB",
		})
		return
	}

	// Validasi tipe file
	allowedMimes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/jpg":  true,
		"image/webp": true,
	}

	openFile, _ := file.Open()
	defer openFile.Close()

	// Baca header untuk deteksi tipe file yang tepat
	header := make([]byte, 512)
	openFile.Read(header)
	mimeType := http.DetectContentType(header)

	if !allowedMimes[mimeType] {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_FILE_TYPE",
			Message: "Hanya file JPEG, PNG, dan WebP yang diizinkan",
		})
		return
	}

	// Buat direktori uploads jika belum ada
	uploadDir := "uploads/obat"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "UPLOAD_ERROR",
			Message: "Gagal membuat direktori upload",
		})
		return
	}

	// Generate filename yang unik
	timestamp := time.Now().Unix()
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("obat_%d%s", timestamp, ext)
	filepath := filepath.Join(uploadDir, filename)

	// Reset file pointer ke awal
	openFile, _ = file.Open()
	defer openFile.Close()

	// Buat file di server
	dst, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "UPLOAD_ERROR",
			Message: "Gagal menyimpan file",
		})
		return
	}
	defer dst.Close()

	// Copy file ke server
	if _, err := io.Copy(dst, openFile); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "UPLOAD_ERROR",
			Message: "Gagal menyimpan file",
		})
		return
	}

	// Return file path with absolute URL
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"filename": filename,
			"path":     filepath,
			"url":      fmt.Sprintf("/uploads/obat/%s", filename),
		},
	})
}

// UploadBase64Image menangani upload image dalam format base64
func (h *FileHandler) UploadBase64Image(c *gin.Context) {
	var req struct {
		Image    string `json:"image" binding:"required"`
		Filename string `json:"filename"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	// TODO: Decode base64 dan simpan
	// Buat direktori uploads jika belum ada
	uploadDir := "uploads/obat"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "UPLOAD_ERROR",
			Message: "Gagal membuat direktori upload",
		})
		return
	}

	// Generate filename yang unik
	timestamp := time.Now().Unix()
	ext := ".jpg" // default untuk base64
	if req.Filename != "" {
		ext = filepath.Ext(req.Filename)
	}
	filename := fmt.Sprintf("obat_%d%s", timestamp, ext)
	filepath := filepath.Join(uploadDir, filename)

	// Simpan base64 image (akan di-decode nantinya)
	// Untuk sekarang, simpan path saja
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"filename": filename,
			"path":     filepath,
			// "url":      fmt.Sprintf("http://localhost:8080/uploads/obat/%s", filename),
			"url":      fmt.Sprintf("/uploads/obat/%s", filename),
		},
	})
}
