package main

import (
	"net/http"
	"preorder/controllers"
	"preorder/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router = controllers.ApplyAuthorRouter(router)
	router = controllers.ApplyFormatRouter(router)
	router = controllers.ApplyOrderRouter(router)

	return router
}

func main() {
	router := SetupRouter()

	models.ConnectDatabase()

	s := &http.Server{
		Addr:    ":8123",
		Handler: router,
	}
	s.ListenAndServe()
}
