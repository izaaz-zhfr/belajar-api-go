package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi/config"
	"goapi/models"
)

// GET /api/siswa - semua siswa + kelas + kartu
func GetSiswa(c *gin.Context) {
	var siswa []models.Siswa
	if err := config.DB.Preload("Kelas").Preload("KartuPelajar").Find(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, siswa)
}

// GET /api/siswa/:id
func GetSiswaByID(c *gin.Context) {
	var s models.Siswa
	if err := config.DB.Preload("Kelas").Preload("KartuPelajar").First(&s, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// POST /api/siswa  body: {"nama": "...", "id_kelas": 1}
func CreateSiswa(c *gin.Context) {
	var input struct {
		Nama    string `json:"nama" binding:"required"`
		IdKelas uint   `json:"id_kelas" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := models.Siswa{Nama: input.Nama, IdKelas: input.IdKelas}
	if err := config.DB.Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	config.DB.Preload("Kelas").First(&s, s.ID)
	c.JSON(http.StatusCreated, s)
}

// PUT /api/siswa/:id
func UpdateSiswa(c *gin.Context) {
	var s models.Siswa
	if err := config.DB.First(&s, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan"})
		return
	}
	var input struct {
		Nama    string `json:"nama"`
		IdKelas uint   `json:"id_kelas"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Nama != "" {
		s.Nama = input.Nama
	}
	if input.IdKelas != 0 {
		s.IdKelas = input.IdKelas
	}
	config.DB.Save(&s)
	c.JSON(http.StatusOK, s)
}

// DELETE /api/siswa/:id
func DeleteSiswa(c *gin.Context) {
	var s models.Siswa
	if err := config.DB.First(&s, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan"})
		return
	}
	config.DB.Delete(&s)
	c.JSON(http.StatusOK, gin.H{"message": "Siswa dihapus"})
}
