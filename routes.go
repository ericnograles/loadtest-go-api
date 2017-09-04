package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initializeRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/hello", hello(db))
	router.GET("/loaderio-c7b3086d5fba14d88dff2c652577b001.txt", func(c *gin.Context) {
		c.String(http.StatusOK, "loaderio-c7b3086d5fba14d88dff2c652577b001")
	})

	return router
}

func hello(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := db.Where("email_address in (?)", []string{"grales@gmail.com"}).Find(&[]User{})
		c.JSON(http.StatusOK, query.Value)
	}
}
