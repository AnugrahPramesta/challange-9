package controllers

import (
	"chal9/database"
	"chal9/helpers"
	"chal9/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	db := database.GetDB()
	product := []models.Product{}
	err := db.Find(&product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, product)
}
func GetProductById(c *gin.Context) {
	db := database.GetDB()
	productId, err := strconv.Atoi(c.Param("productId"))
	product := models.Product{}
	err = db.Model(&product).Where("id = ?", productId).First(&product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "Data Not Found")
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	role := userData["role"].(string)
	if role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Hanya Admin Yang Boleh Mengakses Fitur ini")
		return
	}

	contentType := helpers.GetContentType(c)
	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id = ?", productId).Updates(models.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}
func DeleteProduct(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	role := userData["role"].(string)
	if role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Your Are Not Allowed Access This Feature")
		return
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	product := models.Product{}

	err := db.Where("id = ?", productId).Delete(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "Product Deleted")
}
