
package routes

import (
	controllers "santrikoding/backend-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Route register
	router.POST("/api/register", controllers.Register)

	// Route login
	router.POST("/api/login", controllers.Login)

	// Route untuk mengambil semua data user
	router.GET("/api/users", controllers.FindUsers)

	return router
}