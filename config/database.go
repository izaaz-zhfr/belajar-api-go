package config

import (
	"log"

	"github.com/glebarez/sqlite" // Menggunakan driver Pure Go (bebas CGo)
	"gorm.io/gorm"
	"goapi/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("data.sqlite"), &gorm.Config{
		// Tabel sudah ada, jangan buat FK constraint otomatis
		// biar tidak bentrok dengan schema yang ada
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Gagal konek ke database: ", err)
	}

	// Tabel sudah ada di data.sqlite, jadi TIDAK perlu AutoMigrate.
	// Kalau mau bikin tabel otomatis saat kosong, uncomment ini:
	// _ = database.AutoMigrate(&models.Kelas{}, &models.Siswa{}, &models.KartuPelajar{})
	_ = models.Kelas{}
	_ = database.AutoMigrate(&models.Guru{})

	DB = database
}