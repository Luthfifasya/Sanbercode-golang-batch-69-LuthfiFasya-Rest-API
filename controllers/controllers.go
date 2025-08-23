package controllers

import (
	"fmt"
	"net/http"

	"bioskop-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db
}

func CreateBioskop(c *gin.Context) {
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	bioskop := models.Bioskop{Nama: input.Nama, Lokasi: input.Lokasi, Rating: input.Rating}
	if err := DB.Create(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     bioskop.ID,
		"nama":   bioskop.Nama,
		"lokasi": bioskop.Lokasi,
		"rating": fmt.Sprintf("%.1f", bioskop.Rating),
	})
}

func GetAllBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	if err := DB.Find(&bioskops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	var result []gin.H
	for _, b := range bioskops {
		result = append(result, gin.H{
			"id":     b.ID,
			"nama":   b.Nama,
			"lokasi": b.Lokasi,
			"rating": fmt.Sprintf("%.1f", b.Rating),
		})
	}
	c.JSON(http.StatusOK, result)
}

func GetBioskopByID(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     bioskop.ID,
		"nama":   bioskop.Nama,
		"lokasi": bioskop.Lokasi,
		"rating": fmt.Sprintf("%.1f", bioskop.Rating),
	})
}

func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	bioskop.Nama = input.Nama
	bioskop.Lokasi = input.Lokasi
	bioskop.Rating = input.Rating

	if err := DB.Save(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     bioskop.ID,
		"nama":   bioskop.Nama,
		"lokasi": bioskop.Lokasi,
		"rating": fmt.Sprintf("%.1f", bioskop.Rating),
	})
}

func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&models.Bioskop{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
