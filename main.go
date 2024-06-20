package main

import (
	"gorm.io/gorm"
	"net/http"
	"preorder/authors"
	"preorder/config"
	"preorder/formats"
	"preorder/orders"
	"preorder/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router = authors.ApplyAuthorRouter(router)
	router = formats.ApplyFormatRouter(router)
	router = orders.ApplyOrderRouter(router)

	return router
}

func Migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&authors.Author{})
	_ = db.AutoMigrate(&formats.Format{})
	_ = db.AutoMigrate(&orders.Order{})
	_ = db.AutoMigrate(&users.User{})
}

func main() {
	router := SetupRouter()

	config.ConnectDatabase()
	Migrate(config.DB)

	s := &http.Server{
		Addr:    ":8123",
		Handler: router,
	}
	s.ListenAndServe()
}
