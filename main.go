package main

import (
	"api-hackathon/controllers"
	_ "api-hackathon/cron"
	"api-hackathon/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/login", controllers.Login)
	v1.Use(middleware.AuthRequired)
	{
		v1.GET("/hackathons", controllers.GetHackathones)
		v1.GET("/hackathons/top-ten", controllers.GetDevelopments)
	}

	r.Run()
}
