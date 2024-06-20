package formats

import (
	"net/http"
	"preorder/config"

	"github.com/gin-gonic/gin"
)

func FindFormats(c *gin.Context) {
	var formats []Format
	config.DB.Find(&formats)

	c.JSON(http.StatusOK, gin.H{"data": formats})
}

func FindFormat(c *gin.Context) {
	var format Format

	if err := config.DB.Where("id = ?", c.Param("id")).First(&format).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": format})
}

func CreateFormat(c *gin.Context) {
	var input CreateFormatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	format := NewFormat(input.ID, input.Format)
	config.DB.Create(&format)
	c.JSON(http.StatusOK, gin.H{"data": format})
}

func UpdateFormat(c *gin.Context) {
	var format Format
	if err := config.DB.Where("id = ?", c.Param("id")).First(&format).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input UpdateFormatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&format).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": format})
}

func DeleteFormat(c *gin.Context) {
	var format Format
	if err := config.DB.Where("id = ?", c.Param("id")).First(&format).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&format)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
