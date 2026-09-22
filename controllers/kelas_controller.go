package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi/config"
	"goapi/models"
)

// GET /api/kelas - ambil semua kelas + siswanya
func GetKelas(c *gin.Context) {
	var kelas []models.Kelas
	if err := config.DB.Preload("Siswas").Find(&kelas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kelas)
}

// GET /api/kelas/:id
func GetKelasByID(c *gin.Context) {
	var k models.Kelas
	if err := config.DB.Preload("Siswas").First(&k, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, k)
}

// POST /api/kelas  body: {"nama_kelas": "X RPL 1"}
func CreateKelas(c *gin.Context) {
	var input struct {
		NamaKelas string `json:"nama_kelas" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	k := models.Kelas{NamaKelas: input.NamaKelas}
	if err := config.DB.Create(&k).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, k)
}

// PUT /api/kelas/:id
func UpdateKelas(c *gin.Context) {
	var k models.Kelas
	if err := config.DB.First(&k, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas tidak ditemukan"})
		return
	}
	var input struct {
		NamaKelas string `json:"nama_kelas" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&k).Update("nama_kelas", input.NamaKelas)
	c.JSON(http.StatusOK, k)
}

// DELETE /api/kelas/:id
func DeleteKelas(c *gin.Context) {
	var k models.Kelas
	if err := config.DB.First(&k, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas tidak ditemukan"})
		return
	}
	config.DB.Delete(&k)
	c.JSON(http.StatusOK, gin.H{"message": "Kelas dihapus"})
}
