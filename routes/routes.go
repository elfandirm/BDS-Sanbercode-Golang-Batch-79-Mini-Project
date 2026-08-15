package routes

import (
	"github.com/gin-gonic/gin"

	"formative-13/handlers"
)

func SetupRoutes(router *gin.Engine, bioskopHandler *handlers.BioskopHandler) {
	router.POST("/bioskop", bioskopHandler.CreateBioskop)

	router.GET("/bioskop", bioskopHandler.GetAllBioskop)
	router.GET("/bioskop/:id", bioskopHandler.GetBioskopByID)

	router.PUT("/bioskop/:id", bioskopHandler.UpdateBioskop)

	router.DELETE("/bioskop/:id", bioskopHandler.DeleteBioskop)
}
