package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/techvisionus/golang-source/api/config"
	database "github.com/techvisionus/golang-source/api/db"
	"github.com/techvisionus/golang-source/api/db/gorm/services"
	"github.com/techvisionus/golang-source/api/routes/v1"
	"gorm.io/gorm"
)

type appError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	database.Init()
	// Init connection db
	services.Initialize(config.POSTGRES, &gorm.Config{})
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "server is start")
	})

	v1 := router.Group("/api/v1")
	routes.GetRoutesV1(v1)
	port := config.PORT
	router.Run(":" + port)
}
