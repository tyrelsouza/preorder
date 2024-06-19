package controllers

import (
	"net/http"
	"preorder/models"

	"github.com/gin-gonic/gin"
)

func FindOrders(c *gin.Context) {
	var Orders []models.Order
	models.DB.Find(&Orders)

	c.JSON(http.StatusOK, gin.H{"data": Orders})
}

func FindOrder(c *gin.Context) {
	var Order models.Order

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": Order})
}

func CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.NewOrder(input.ID, input.Title, input.Author, input.Format, input.ISBN13, input.ReleaseDate)
	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func UpdateOrder(c *gin.Context) {
	var Order models.Order
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input models.UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Order).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": Order})
}

func DeleteOrder(c *gin.Context) {
	var Order models.Order
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Order)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ApplyOrderRouter(router *gin.Engine) *gin.Engine {
	router.GET("/orders", FindOrders)
	router.POST("/orders", CreateOrder)
	router.GET("/orders/:id", FindOrder)
	router.PATCH("/orders/:id", UpdateOrder)
	router.DELETE("/orders/:id", DeleteOrder)
	return router
}
