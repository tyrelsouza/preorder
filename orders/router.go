package orders

import "github.com/gin-gonic/gin"

func ApplyOrderRouter(router *gin.Engine) *gin.Engine {
	router.GET("/orders", FindOrders)
	router.GET("/orders/:id/cover", FindOrderCover)
	router.POST("/orders", CreateOrder)
	router.GET("/orders/:id", FindOrder)
	router.PATCH("/orders/:id", UpdateOrder)
	router.DELETE("/orders/:id", DeleteOrder)
	return router
}
