package controllers

import (
	"net/http"

	"bioskop-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func CreateBioskop(c *gin.Context) {
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func GetBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	DB.Find(&bioskops)
	c.JSON(http.StatusOK, bioskops)
}

func GetBioskopByID(c *gin.Context) {
	var bioskop models.Bioskop
	if err := DB.First(&bioskop, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, bioskop)
}

func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	if err := DB.First(&bioskop, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	DB.Model(&bioskop).Updates(input)
	c.JSON(http.StatusOK, bioskop)
}

func DeleteBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	if err := DB.First(&bioskop, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	DB.Delete(&bioskop)
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
