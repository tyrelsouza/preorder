package controllers

import (
	"net/http"
	"preorder/models"

	"github.com/gin-gonic/gin"
)

func FindFormats(c *gin.Context) {
	var Formats []models.Format
	models.DB.Find(&Formats)

	c.JSON(http.StatusOK, gin.H{"data": Formats})
}

func FindFormat(c *gin.Context) {
	var Format models.Format

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Format).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": Format})
}

func CreateFormat(c *gin.Context) {
	var input models.CreateFormatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	format := models.NewFormat(input.ID, input.Format)
	models.DB.Create(&format)
	c.JSON(http.StatusOK, gin.H{"data": format})
}

func UpdateFormat(c *gin.Context) {
	var Format models.Format
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Format).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input models.UpdateFormatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Format).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": Format})
}

func DeleteFormat(c *gin.Context) {
	var Format models.Format
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Format).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Format)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ApplyFormatRouter(router *gin.Engine) *gin.Engine {
	router.GET("/formats", FindFormats)
	router.POST("/formats", CreateFormat)
	router.GET("/formats/:id", FindFormat)
	router.PATCH("/formats/:id", UpdateFormat)
	router.DELETE("/formats/:id", DeleteFormat)
	return router
}
