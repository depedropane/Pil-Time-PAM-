package http

import (
	"backend/internal/dto"
	"backend/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ==================== TrackingJadwal Handler ====================

type TrackingJadwalHandler struct {
	usecase *usecase.TrackingJadwalUsecase
}

func NewTrackingJadwalHandler(u *usecase.TrackingJadwalUsecase) *TrackingJadwalHandler {
	return &TrackingJadwalHandler{u}
}

func (h *TrackingJadwalHandler) GetAll(c *gin.Context) {
	trackings, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "FETCH_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": trackings,
	})
}

func (h *TrackingJadwalHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "ID harus berupa angka",
		})
		return
	}

	tracking, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Tracking jadwal tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tracking,
	})
}

func (h *TrackingJadwalHandler) GetByPasienID(c *gin.Context) {
	pasienID, err := strconv.Atoi(c.Param("pasien_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Pasien ID harus berupa angka",
		})
		return
	}

	trackings, err := h.usecase.GetByPasienID(pasienID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "FETCH_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": trackings,
	})
}

// GetObatStreak - Mengambil jumlah streak kepatuhan meminum obat pasien
func (h *TrackingJadwalHandler) GetObatStreak(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("pasien_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "ID pasien tidak valid",
		})
		return
	}

	s, err := h.usecase.GetObatStreak(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "FETCH_ERROR",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"current_streak": s})
}

// GetMyRiwayat digunakan oleh pasien untuk melihat riwayatnya sendiri.
// pasien_id diambil dari JWT context (inject oleh JWTPasienMiddleware).
func (h *TrackingJadwalHandler) GetMyRiwayat(c *gin.Context) {
	pasienIDRaw, exists := c.Get("pasien_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "UNAUTHORIZED",
			Message: "Sesi tidak valid, silakan login ulang",
		})
		return
	}

	pasienID, ok := pasienIDRaw.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Gagal membaca ID pasien dari token",
		})
		return
	}

	trackings, err := h.usecase.GetByPasienIDIncludingMandiri(pasienID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "FETCH_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": trackings,
	})
}

// CreateMyRiwayat digunakan oleh pasien untuk mencatat riwayat minum obatnya.
// pasien_id diambil dari token (JWT) agar pasien tidak bisa memalsukan ID pasien lain.
func (h *TrackingJadwalHandler) CreateMyRiwayat(c *gin.Context) {
	pasienIDRaw, exists := c.Get("pasien_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "UNAUTHORIZED",
			Message: "Sesi tidak valid, silakan login ulang",
		})
		return
	}

	pasienID := 0
	switch v := pasienIDRaw.(type) {
	case float64:
		pasienID = int(v)
	case int:
		pasienID = v
	case uint:
		pasienID = int(v)
	}

	var req dto.CreateTrackingJadwalDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	// Paksa ID pasien pada payload agar sesuai dengan user yang login
	req.PasienID = pasienID

	result, err := h.usecase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "CREATE_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func (h *TrackingJadwalHandler) Create(c *gin.Context) {
	var req dto.CreateTrackingJadwalDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	if req.PasienID == 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: "Key: 'CreateTrackingJadwalDTO.PasienID' Error:Field validation for 'PasienID' failed on the 'required' tag",
		})
		return
	}

	result, err := h.usecase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "CREATE_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func (h *TrackingJadwalHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "ID harus berupa angka",
		})
		return
	}

	var req dto.UpdateTrackingJadwalDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	result, err := h.usecase.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "UPDATE_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *TrackingJadwalHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "ID harus berupa angka",
		})
		return
	}

	err = h.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "DELETE_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tracking jadwal berhasil dihapus",
	})
}

func (h *TrackingJadwalHandler) GetStatistics(c *gin.Context) {
	pasienID := 0
	if pasienIDStr := c.Query("pasien_id"); pasienIDStr != "" {
		if id, err := strconv.Atoi(pasienIDStr); err == nil {
			pasienID = id
		}
	}

	stats, err := h.usecase.GetStatistics(pasienID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "STATISTICS_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stats,
	})
}
