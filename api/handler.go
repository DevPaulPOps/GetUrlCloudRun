package api

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"microservice.com/appfyt/gcp"
	"net/http"
)

func getURLByName(c *gin.Context) {
	name := c.Param("name")
	c.IndentedJSON(http.StatusOK, gcp.GetURLByName(gcp.ListServicesRequest(gcp.ConnectToGCP(context.Background()), context.Background()), name))
}

func getAllUrlAndName(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gcp.GetAllUrl(gcp.ListServicesRequest(gcp.ConnectToGCP(context.Background()), context.Background())))

}

func LauncServ() {
	r := gin.Default()
	r.Use((cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})))

	r.GET("/service/:name", getURLByName)
	r.GET("/service", getAllUrlAndName)
	r.Run(":3000")
}
