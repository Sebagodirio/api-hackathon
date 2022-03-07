package routes

import (
	"api-hackathon/controllers"
	"api-hackathon/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.POST("/login", controllers.Login)
	v1.Use(middleware.AuthRequired)
	{
		v1.GET("/hackathons", controllers.GetHackathons)
		v1.GET("/hackathons/top-ten", controllers.GetDevelopments)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
