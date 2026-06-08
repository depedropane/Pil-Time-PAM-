package main

import (
	"backend/config"
	inboundHttp "backend/internal/adapters/inbound/http"
	"backend/internal/adapters/outbound/persistence"
	"backend/internal/domain"
	"backend/internal/usecase"
	"backend/pkg/fcm"
	"backend/pkg/middleware"
	"backend/pkg/whatsapp"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. DATABASE INIT
	db := config.InitPostgres()

	db.AutoMigrate(
		&domain.ResepObat{},
		&domain.JadwalObat{},
		&domain.FcmToken{}, // FCM device tokens
	)
	if db == nil {
		log.Fatal("Gagal mengoneksikan database")
	}

	// Init Firebase Admin SDK untuk FCM
	if err := fcm.Init("serviceAccountKey.json"); err != nil {
		log.Printf("[WARNING] Firebase init gagal: %v — FCM notifikasi tidak akan berfungsi", err)
	}

	// 2. DEPENDENCY INJECTION (DI)
	// Repositories
	nakesRepo := persistence.NewNakesRepo(db)
	pasienRepo := persistence.NewPasienRepo(db)
	jadwalRepo := persistence.NewJadwalRepo(db)
	obatRepo := persistence.NewObatRepo(db)
	trackingRepo := persistence.NewTrackingJadwalRepo(db)
	rutinitasRepo := persistence.NewRutinitasRepo(db)
	resepRepo := persistence.NewResepObatRepo(db)
	jadwalObatRepo := persistence.NewJadwalObatRepo(db)
	fcmTokenRepo := persistence.NewFcmTokenRepo(db)

	// Usecases
	adminUC := usecase.NewAdminUsecase(nakesRepo)
	pasienUC := usecase.NewPasienUsecase(pasienRepo, jadwalRepo)
	jadwalUC := usecase.NewJadwalUsecase(jadwalRepo, pasienRepo, fcmTokenRepo)
	dashboardUC := usecase.NewDashboardUsecase(pasienRepo, jadwalRepo)
	obatUC := usecase.NewObatUsecase(obatRepo, jadwalRepo)
	trackingUC := usecase.NewTrackingJadwalUsecase(trackingRepo, jadwalRepo, pasienRepo)
	rutinitasUC := usecase.NewRutinitasUsecase(rutinitasRepo)
	resepJadwalUC := usecase.NewResepJadwalUsecase(resepRepo, jadwalObatRepo, jadwalRepo, obatRepo)
	authUC := usecase.NewAuthUsecase(nakesRepo, pasienRepo)

	// Handlers
	adminHandler := inboundHttp.NewAdminHandler(adminUC)
	pasienHandler := inboundHttp.NewPasienHandler(pasienUC)
	jadwalHandler := inboundHttp.NewJadwalHandler(jadwalUC)
	dashboardHandler := inboundHttp.NewDashboardHandler(dashboardUC)
	obatHandler := inboundHttp.NewObatHandler(obatUC)
	trackingHandler := inboundHttp.NewTrackingJadwalHandler(trackingUC)
	rutinitasHandler := inboundHttp.NewRutinitasHandler(rutinitasUC)
	resepJadwalHandler := inboundHttp.NewResepJadwalHandler(resepJadwalUC)
	fileHandler := inboundHttp.NewFileHandler()
	fcmTokenHandler := inboundHttp.NewFcmTokenHandler(fcmTokenRepo)
	authHandler := inboundHttp.NewAuthHandler(authUC)

	// 3. ROUTER SETUP
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(CORSConfig())

	// Unified Auth Service routes (monolith)
	auth := r.Group("/auth")
	{
		nakes := auth.Group("/nakes")
		{
			nakes.POST("/login", authHandler.LoginNakes)
		}

		pasien := auth.Group("/pasien")
		{
			pasien.POST("/register", authHandler.RegisterPasien)
			pasien.POST("/login", authHandler.LoginPasien)
			pasien.POST("/forgot-password", authHandler.ForgotPassword)
			pasien.POST("/verify-reset-code", authHandler.VerifyResetCode)
			pasien.POST("/reset-password", authHandler.ResetPassword)
		}

		auth.GET("/validate", authHandler.ValidateToken)
	}

	// 4. STATIC FILES (PENTING untuk PWA & Uploads)
	r.Static("/uploads", "./public/uploads")
	r.StaticFile("/manifest.json", "./public/manifest.json")
	r.Static("/img", "./public/img")

	// 5. API ROUTES
	api := r.Group("/api")
	{
		// Endpoint untuk test WA langsung dari browser (tanpa auth)
		api.GET("/test-wa", func(c *gin.Context) {
			phone := c.Query("phone")
			if phone == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "phone parameter is required (e.g. /api/test-wa?phone=08123456789)"})
				return
			}
			// Memanggil fungsi SendWarning yang sudah ada di package whatsapp
			whatsapp.SendWarning(phone, "Pasien Uji Coba", "Obat Paracetamol", time.Now().Format("15:04"))
			c.JSON(http.StatusOK, gin.H{
				"message": "Pesan peringatan berhasil dikirimkan ke " + phone,
				"info": "Cek log terminal untuk detail response dari Fonnte",
			})
		})

		// --- ADMIN ROUTES ---
		admin := api.Group("/admin")
		{
			// Login tidak butuh middleware JWT
			admin.POST("/login", adminHandler.Login)

			// Semua route admin lainnya diproteksi JWT nakes
			adminProtected := admin.Group("")
			adminProtected.Use(middleware.JWTNakesMiddleware())
			{
				adminProtected.GET("/dashboard", dashboardHandler.GetDashboard)
				adminProtected.GET("/pasien", pasienHandler.GetAll)

				// Admin - Info Obat
				obat := adminProtected.Group("/info-obat")
				{
					obat.GET("", obatHandler.GetAll)
					obat.GET("/:id", obatHandler.GetByID)
					obat.POST("", obatHandler.Create)
					obat.PUT("/:id", obatHandler.Update)
					obat.DELETE("/:id", obatHandler.Delete)
				}

				// Admin - Jadwal
				jadwal := adminProtected.Group("/jadwal")
				{
					jadwal.GET("", jadwalHandler.GetAllJadwal)
					jadwal.GET("/:id", jadwalHandler.GetJadwalByID)
					jadwal.GET("/pasien/:pasien_id", jadwalHandler.GetJadwalByPasien)
					jadwal.POST("", jadwalHandler.CreateJadwal)
					jadwal.PUT("/:id", jadwalHandler.UpdateJadwal)
					jadwal.DELETE("/:id", jadwalHandler.DeleteJadwal)
				}
				adminProtected.POST("/resep-jadwal", resepJadwalHandler.Create)

				// Admin - Tracking/Riwayat
				riwayat := adminProtected.Group("/riwayat")
				{
					riwayat.GET("", trackingHandler.GetAll)
					riwayat.GET("/statistics", trackingHandler.GetStatistics)
					riwayat.GET("/:id", trackingHandler.GetByID)
					riwayat.GET("/pasien/:pasien_id", trackingHandler.GetByPasienID)
					riwayat.POST("", trackingHandler.Create)
					riwayat.PUT("/:id", trackingHandler.Update)
					riwayat.DELETE("/:id", trackingHandler.Delete)
				}
			}
		}

		// --- PASIEN AUTHENTICATED ROUTES ---
		// Semua route pasien diproteksi JWT pasien (dari auth-service)
		pasienAuth := api.Group("/pasien")
		pasienAuth.Use(middleware.JWTPasienMiddleware())
		{
			pasienAuth.GET("/dashboard", pasienHandler.GetDashboard)
			pasienAuth.GET("/jadwal", pasienHandler.GetJadwal)
			pasienAuth.GET("/profile", pasienHandler.GetProfile)
			pasienAuth.PUT("/profile", pasienHandler.UpdateProfile)
			pasienAuth.POST("/profile/test-wa-warning", pasienHandler.TestWaWarning)

			// Pasien - Rutinitas
			pasienAuth.GET("/rutinitas", rutinitasHandler.GetAllForPasien)
			pasienAuth.GET("/rutinitas/streak/:pasien_id", rutinitasHandler.GetStreak)
			pasienAuth.POST("/rutinitas/tracking", rutinitasHandler.UpdateTracking)
			pasienAuth.POST("/rutinitas", rutinitasHandler.Create)
			pasienAuth.PUT("/rutinitas/:id", rutinitasHandler.Update)
			pasienAuth.DELETE("/rutinitas/:id", rutinitasHandler.Delete)

			// Pasien - Riwayat Kepatuhan (pasien melihat/mencatat riwayatnya sendiri)
			pasienAuth.GET("/riwayat", trackingHandler.GetMyRiwayat)
			pasienAuth.GET("/riwayat/streak/:pasien_id", trackingHandler.GetObatStreak)
			pasienAuth.POST("/riwayat", trackingHandler.CreateMyRiwayat)

			// Pasien - FCM Token Registration
			pasienAuth.POST("/fcm-token", fcmTokenHandler.RegisterToken)

			// Pasien - Obat Mandiri
			pasienAuth.POST("/obat-mandiri", obatHandler.CreateMandiri)
			pasienAuth.GET("/obat-mandiri", obatHandler.GetAllForPasien)
			pasienAuth.GET("/obat-mandiri/:id", obatHandler.GetByID)
			pasienAuth.PUT("/obat-mandiri/:id", obatHandler.Update)
			pasienAuth.DELETE("/obat-mandiri/:id", obatHandler.Delete)
		}

		// --- UPLOAD ROUTES ---
		api.POST("/upload/image", fileHandler.UploadImage)
		api.POST("/upload/image-base64", fileHandler.UploadBase64Image)
	}

	// Start background workers
	warningWorker := usecase.NewWaWarningWorker(db)
	warningWorker.Start()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Pil Time Server running on :%s\n", port)
	r.Run(":" + port)
}

// CORSConfig mengatur CORS untuk semua request
func CORSConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
