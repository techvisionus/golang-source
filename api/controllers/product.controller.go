package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	constants "github.com/techvisionus/golang-source/api/contants"
	database "github.com/techvisionus/golang-source/api/db"
	dbSchemas "github.com/techvisionus/golang-source/api/db/sqlc"
	utils "github.com/techvisionus/golang-source/api/utils"
	validators "github.com/techvisionus/golang-source/api/validators"
)

// AddProductController controller
func AddProductController(ctx * gin.Context) {
	var req constants.ProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err) )
		return
	}

	productParam := dbSchemas.CreateProductParams {
		Code: req.Code,
		Name: req.Name,
		Price: req.Price,
	}
	valid := validators.ValidCreateProduct(ctx, productParam.Code)

	if !valid {
		return
	}

	product, err := database.Queries.CreateProduct(ctx, productParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err) )
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "Add product success.",
		"data": product,
	})
}
// GetProductController controller
func GetProductController(ctx * gin.Context){
	var req constants.GetProductRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err) )
		return
	}

	product, err := database.Queries.GetProduct(ctx, req.ID)

	if err != nil {
		errMessage := fmt.Errorf("Product not found")
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errMessage) )
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "Get product success",
		"data": product,
	})
}

