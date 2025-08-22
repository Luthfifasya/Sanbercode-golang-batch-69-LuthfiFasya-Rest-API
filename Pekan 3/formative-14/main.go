package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Bioskop struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=lightyagami321 dbname=bioskop_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database")
	}

	db.AutoMigrate(&Bioskop{})
}

func main() {
	initDB()

	r := gin.Default()

	r.POST("/bioskop", func(c *gin.Context) {
		var input Bioskop

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//validasi
		if input.Nama == "" || input.Lokasi == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
			return
		}

		bioskop := Bioskop{
			Nama:   input.Nama,
			Lokasi: input.Lokasi,
			Rating: input.Rating,
		}

		if err := db.Create(&bioskop).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
			return
		}

		c.JSON(http.StatusOK, bioskop)
	})

	r.GET("/bioskop", func(c *gin.Context) {
		var bioskops []Bioskop
		if err := db.Find(&bioskops).Error; err != nil {
			c.JSON(500, gin.H{"error": "Gagal mengambil data"})
			return
		}
		c.JSON(200, bioskops)
	})

	r.Run(":8080")
}
