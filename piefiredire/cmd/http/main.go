package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/adapter/handlers/api"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/adapter/storage/repository"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
	}
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}

	router.Use(cors.New(config))

	router.GET("/checkhealth", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	var mockDB map[string]interface{}

	beefRepository := repository.NewBeefRepository(mockDB)
	beefHandler := api.NewBeefHandler(beefRepository)

	router.GET("/beef/summary", api.GinHandler(beefHandler.GetSummary))

	router.Run()
}
