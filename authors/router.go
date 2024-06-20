package authors

import "github.com/gin-gonic/gin"

func ApplyAuthorRouter(router *gin.Engine) *gin.Engine {
	router.GET("/authors", FindAuthors)
	router.POST("/authors", CreateAuthor)
	router.GET("/authors/:id", FindAuthor)
	router.PATCH("/authors/:id", UpdateAuthor)
	router.DELETE("/authors/:id", DeleteAuthor)
	return router
}
