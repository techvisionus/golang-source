package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/techvisionus/golang-source/api/controllers/v1"
)

func addProductRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/products")

	product.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, "users") })
	// product.GET("/:id", controllers.GetProductController)
	// product.POST("/", controllers.AddProductController)
	product.GET("/:id", controllers.GetProduct)
	product.POST("/", controllers.CreateProduct)
	product.PUT("/", controllers.UpdateProduct)
	product.DELETE("/:id", controllers.DeleteProduct)
}
