package main

import (
	"log"
	"os"

	"bioskop-app/controllers"
	"bioskop-app/models"
	"bioskop-app/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL tidak ditemukan, cek Railway Variables")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	DB = db
	DB.AutoMigrate(&models.Bioskop{})
}

func main() {
	initDB()
	controllers.SetDB(DB)

	r := gin.Default()
	routers.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
