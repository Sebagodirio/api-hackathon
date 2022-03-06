package main

import (
	"api-hackathon/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/hackathones", controllers.GetHackathones)
		v1.GET("/hackathones/top-ten", controllers.GetDevelopments)
	}

	r.Run()
}
