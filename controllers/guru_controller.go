package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi/config"
	"goapi/models"
)

// GET /api/guru - semua guru
func GetGuru(c *gin.Context) {
	var guru []models.Guru
	if err := config.DB.Find(&guru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guru)
}

// GET /api/guru/:id
func GetGuruByID(c *gin.Context) {
	var g models.Guru
	if err := config.DB.First(&g, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, g)
}

// POST /api/guru  body: {"nama": "...", "mapel": "..."}
func CreateGuru(c *gin.Context) {
	var input struct {
		Nama  string `json:"nama" binding:"required"`
		Mapel string `json:"mapel" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g := models.Guru{Nama: input.Nama, Mapel: input.Mapel}
	if err := config.DB.Create(&g).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, g)
}

// PUT /api/guru/:id
func UpdateGuru(c *gin.Context) {
	var g models.Guru
	if err := config.DB.First(&g, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru tidak ditemukan"})
		return
	}
	var input struct {
		Nama  string `json:"nama"`
		Mapel string `json:"mapel"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Nama != "" {
		g.Nama = input.Nama
	}
	if input.Mapel != "" {
		g.Mapel = input.Mapel
	}
	config.DB.Save(&g)
	c.JSON(http.StatusOK, g)
}

// DELETE /api/guru/:id
func DeleteGuru(c *gin.Context) {
	var g models.Guru
	if err := config.DB.First(&g, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru tidak ditemukan"})
		return
	}
	config.DB.Delete(&g)
	c.JSON(http.StatusOK, gin.H{"message": "Guru dihapus"})
}