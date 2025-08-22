package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bioskop-app/controllers"
	"bioskop-app/models"
	"bioskop-app/routers"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=YOUR_PASSWORD dbname=bioskop_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}
	fmt.Println("Berhasil konek ke database!")

	db.AutoMigrate(&models.Bioskop{})
	return db
}

func main() {
	db := initDB()
	controllers.SetDatabase(db)

	r := gin.Default()
	routers.BioskopRouter(r)
	r.Run(":8080")
}
