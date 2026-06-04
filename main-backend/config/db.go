package config

import (
	"backend/internal/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database!")
	}

	log.Println("Database connected!")

	// Auto-migrate schema (Menambahkan Rutinitas, Tracking, dan Riwayat ke daftar migrasi)
	if err := db.AutoMigrate(
		&domain.Pasien{}, 
		&domain.Nakes{}, 
		&domain.Jadwal{}, 
		&domain.Obat{},
		&domain.Rutinitas{},         // <--- Tambahan
		&domain.TrackingRutinitas{}, // <--- Tambahan
		&domain.TrackingJadwal{},    // <--- Tambahan
		&domain.RiwayatObat{},       // <--- Tambahan
		&domain.WaWarning{},         // <--- Tambahan
	); err != nil {
		panic("Gagal auto-migrate: " + err.Error())
	}
	log.Println("Database migration completed!")

	return db
}