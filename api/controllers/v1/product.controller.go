package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/techvisionus/golang-source/api/db/gorm/models"
	"github.com/techvisionus/golang-source/api/db/gorm/services"
)

var CreateProduct = func(ctx *gin.Context) {
	log.Info("Begin - CreateProduct")
	db := services.GetDB()
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if nil != err {
		log.Error("Can not get data from requestBody. err=", err)
		return
	}
	if err = db.Create(&product).Error; nil != err {
		log.Error("Can not save product. err=", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Add product success.",
		"data":    product,
	})
	log.Info("End - CreateProduct")
}

var GetProduct = func(ctx *gin.Context) {
	log.Info("Begin - GetProduct")
	db := services.GetDB()
	var product models.Product
	id := ctx.Param("id")
	if len(id) < 0 {
		log.Error("Can not get ID.")
		return
	}
	if err := db.Where("id = ?", id).First(&product).Error; nil != err {
		log.Error("Can not get product. err=", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product success.",
		"data":    product,
	})
	log.Info("End - GetProduct")
}

var UpdateProduct = func(ctx *gin.Context) {
	log.Info("Begin - UpdateProduct")
	db := services.GetDB()
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if nil != err {
		log.Error("Can not get data from requestBody. err=", err)
		return
	}
	if err = db.Save(&product).Error; nil != err {
		log.Error("Can not update product. err=", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update product success.",
		"data":    product,
	})
	log.Info("End - UpdateProduct")
}

var DeleteProduct = func(ctx *gin.Context) {
	log.Info("Begin - DeleteProduct")
	db := services.GetDB()
	id := ctx.Param("id")
	if len(id) < 0 {
		log.Error("Can not get ID.")
		return
	}
	if err := db.Where("id = ?", id).Delete(&models.Product{}).Error; nil != err {
		log.Error("Can not delete product. err=", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete product success.",
		"data":    nil,
	})
	log.Info("End - DeleteProduct")
}
