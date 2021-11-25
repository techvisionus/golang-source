package main

import (
	"net/http"
	"os"

	database "github.com/techvisionus/golang-source/api/db"
	"github.com/techvisionus/golang-source/api/routes/v1"

	"github.com/gin-gonic/gin"
)

type appError struct {
    Code     int    `json:"code"`
    Message  string `json:"message"`
}

func main() {
	router := gin.Default()
	database.Init()

	router.GET("/", func (context *gin.Context) {
		context.String(http.StatusOK, "server is start")
	})

	v1 := router.Group("/api/v1")
	routes.GetRoutesV1(v1)
	port := os.Getenv("PORT")
	router.Run(":"+ port)
}



