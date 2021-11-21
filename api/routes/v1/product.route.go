package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/techvisionus/golang-source/api/controllers"
)

func addProductRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/products")

	product.GET("/", func(c *gin.Context) {c.JSON(http.StatusOK, "users")})
	product.GET("/:id", controllers.GetProductController)
	product.POST("/", controllers.AddProductController)
}