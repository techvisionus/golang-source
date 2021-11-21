package routes

import "github.com/gin-gonic/gin"

// GetRoutesV1 routes
func GetRoutesV1(rg *gin.RouterGroup) {
	addPingRoutes(rg)
	addProductRoutes(rg)
}