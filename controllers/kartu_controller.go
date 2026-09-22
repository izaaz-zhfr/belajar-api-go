package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi/config"
	"goapi/models"
)

// GET /api/kartu - semua kartu + data siswa
func GetKartu(c *gin.Context) {
	var kartu []models.KartuPelajar
	if err := config.DB.Preload("Siswa.Kelas").Find(&kartu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kartu)
}

// GET /api/kartu/:id
func GetKartuByID(c *gin.Context) {
	var k models.KartuPelajar
	if err := config.DB.Preload("Siswa.Kelas").First(&k, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kartu tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, k)
}

// POST /api/kartu  body: {"nomor_kartu": "KP-2026-000001", "id_siswa": 1}
func CreateKartu(c *gin.Context) {
	var input struct {
		NomorKartu string `json:"nomor_kartu" binding:"required"`
		IdSiswa    uint   `json:"id_siswa" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	k := models.KartuPelajar{NomorKartu: input.NomorKartu, IdSiswa: input.IdSiswa}
	if err := config.DB.Create(&k).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, k)
}

// PUT /api/kartu/:id
func UpdateKartu(c *gin.Context) {
	var k models.KartuPelajar
	if err := config.DB.First(&k, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kartu tidak ditemukan"})
		return
	}
	var input struct {
		NomorKartu string `json:"nomor_kartu"`
		IdSiswa    uint   `json:"id_siswa"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.NomorKartu != "" {
		k.NomorKartu = input.NomorKartu
	}
	if input.IdSiswa != 0 {
		k.IdSiswa = input.IdSiswa
	}
	if err := config.DB.Save(&k).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, k)
}

// DELETE /api/kartu/:id
func DeleteKartu(c *gin.Context) {
	var k models.KartuPelajar
	if err := config.DB.First(&k, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kartu tidak ditemukan"})
		return
	}
	config.DB.Delete(&k)
	c.JSON(http.StatusOK, gin.H{"message": "Kartu dihapus"})
}
