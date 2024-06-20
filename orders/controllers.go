package orders

import (
	"net/http"
	"preorder/config"

	"github.com/gin-gonic/gin"
)

func FindOrders(c *gin.Context) {
	var Orders []Order
	config.DB.Find(&Orders)

	c.JSON(http.StatusOK, gin.H{"data": Orders})
}

func FindOrder(c *gin.Context) {
	var order Order

	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := NewOrder(input.ID, input.Title, input.Author, input.Format, input.ISBN13, input.ReleaseDate)
	config.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func UpdateOrder(c *gin.Context) {
	var order Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&order).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrder(c *gin.Context) {
	var order Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
