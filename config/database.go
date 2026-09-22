package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"goapi/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("data.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database: ", err)
	}

	// Auto migrate sesuai tabel yang sudah ada
	err = database.AutoMigrate(&models.Kelas{}, &models.Siswa{}, &models.KartuPelajar{})
	if err != nil {
		log.Fatal("Gagal migrate: ", err)
	}

	DB = database
}
