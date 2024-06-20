package authors

import (
	"net/http"
	"preorder/config"

	"github.com/gin-gonic/gin"
)

func FindAuthors(c *gin.Context) {
	var authors []Author
	config.DB.Find(&authors)

	c.JSON(http.StatusOK, gin.H{"data": authors})
}

func FindAuthor(c *gin.Context) {
	var author Author

	if err := config.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func CreateAuthor(c *gin.Context) {
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := NewAuthor(input.ID, input.FullName)
	config.DB.Create(&author)
	c.JSON(http.StatusOK, gin.H{"data": author})
}

func UpdateAuthor(c *gin.Context) {
	var author Author
	if err := config.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input UpdateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&author).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": author})
}

func DeleteAuthor(c *gin.Context) {
	var author Author
	if err := config.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&author)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
