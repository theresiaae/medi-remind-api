package routes

import (
	"medi_remind_backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Selamat datang di Medi Remind!"})
	})

	// Tambahkan route register & login
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/users", controllers.GetAllUsers)


	obatRoutes := r.Group("/obat")
	{
		obatRoutes.POST("/", controllers.CreateObat)
		obatRoutes.GET("/", controllers.GetAllObat)
		obatRoutes.GET("/:id", controllers.GetObatByID)
		obatRoutes.PUT("/:id", controllers.UpdateObat)
		obatRoutes.DELETE("/:id", controllers.DeleteObat)
	}

}
