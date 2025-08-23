package routers

import (
	"bioskop-app/controllers"

	"github.com/gin-gonic/gin"
)

func BioskopRouter(r *gin.Engine) {
	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetAllBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopByID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)
}
