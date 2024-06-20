package formats

import "github.com/gin-gonic/gin"

func ApplyFormatRouter(router *gin.Engine) *gin.Engine {
	router.GET("/formats", FindFormats)
	router.POST("/formats", CreateFormat)
	router.GET("/formats/:id", FindFormat)
	router.PATCH("/formats/:id", UpdateFormat)
	router.DELETE("/formats/:id", DeleteFormat)
	return router
}
