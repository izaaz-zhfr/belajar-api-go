package main

import (
	"github.com/gin-gonic/gin"
	"goapi/config"
	"goapi/controllers"
)

func main() {
	// Konek ke SQLite (data.sqlite)
	config.ConnectDatabase()

	r := gin.Default()

	// Cek server hidup
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Sekolah jalan! Coba /api/kelas, /api/siswa, /api/kartu, /api/guru"})
	})

	api := r.Group("/api")
	{
		// Kelas
		api.GET("/kelas", controllers.GetKelas)
		api.GET("/kelas/:id", controllers.GetKelasByID)
		api.POST("/kelas", controllers.CreateKelas)
		api.PUT("/kelas/:id", controllers.UpdateKelas)
		api.DELETE("/kelas/:id", controllers.DeleteKelas)

		// Siswa
		api.GET("/siswa", controllers.GetSiswa)
		api.GET("/siswa/:id", controllers.GetSiswaByID)
		api.POST("/siswa", controllers.CreateSiswa)
		api.PUT("/siswa/:id", controllers.UpdateSiswa)
		api.DELETE("/siswa/:id", controllers.DeleteSiswa)

		// Kartu Pelajar
		api.GET("/kartu", controllers.GetKartu)
		api.GET("/kartu/:id", controllers.GetKartuByID)
		api.POST("/kartu", controllers.CreateKartu)
		api.PUT("/kartu/:id", controllers.UpdateKartu)
		api.DELETE("/kartu/:id", controllers.DeleteKartu)

		// Guru
		api.GET("/guru", controllers.GetGuru)
		api.GET("/guru/:id", controllers.GetGuruByID)
		api.POST("/guru", controllers.CreateGuru)
		api.PUT("/guru/:id", controllers.UpdateGuru)
		api.DELETE("/guru/:id", controllers.DeleteGuru)
	}

	r.Run(":8080") // jalan di http://localhost:8080
}
