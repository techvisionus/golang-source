package validators

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/techvisionus/golang-source/api/db"
	utils "github.com/techvisionus/golang-source/api/utils"
)

// ValidCreateProduct validate product
func ValidCreateProduct(ctx *gin.Context, productCode string) bool {
	_, err := database.Queries.GetProductByCode(ctx, productCode)
	if err == nil {
		errMessage := fmt.Errorf("Code %s already exists", productCode)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errMessage))
		return false
	}
	return true
}