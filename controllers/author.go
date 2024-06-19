package controllers

import (
	"net/http"
	"preorder/models"

	"github.com/gin-gonic/gin"
)

func FindAuthors(c *gin.Context) {
	var Authors []models.Author
	models.DB.Find(&Authors)

	c.JSON(http.StatusOK, gin.H{"data": Authors})
}

func FindAuthor(c *gin.Context) {
	var Author models.Author

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": Author})
}

func CreateAuthor(c *gin.Context) {
	var input models.CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.NewAuthor(input.ID, input.FullName)
	models.DB.Create(&author)
	c.JSON(http.StatusOK, gin.H{"data": author})
}

func UpdateAuthor(c *gin.Context) {
	var Author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input models.UpdateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Author).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": Author})
}

func DeleteAuthor(c *gin.Context) {
	var Author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Author)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ApplyAuthorRouter(router *gin.Engine) *gin.Engine {
	router.GET("/authors", FindAuthors)
	router.POST("/authors", CreateAuthor)
	router.GET("/authors/:id", FindAuthor)
	router.PATCH("/authors/:id", UpdateAuthor)
	router.DELETE("/authors/:id", DeleteAuthor)
	return router
}
