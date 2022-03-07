package main

import (
	_ "api-hackathon/cron"
	"api-hackathon/docs"
	"api-hackathon/routes"

	"github.com/gin-gonic/gin"
)

// @title api-hackathon
// @version 2.0
// @description API Restful for getting the best developments in the hackathons
// @termsOfService http://swagger.io/terms/

// @contact.name Sebastian Godirio
// @contact.email sebastianlgodirio@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes.SetupRouter(r)
	r.Run()
}
