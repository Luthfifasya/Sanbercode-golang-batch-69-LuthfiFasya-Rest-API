package routers

import (
	"bioskop-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/", controllers.CreateBioskop)
	r.GET("/", controllers.GetBioskop)
	r.GET("/:id", controllers.GetBioskopByID)
	r.PUT("/:id", controllers.UpdateBioskop)
	r.DELETE(":id", controllers.DeleteBioskop)
}
