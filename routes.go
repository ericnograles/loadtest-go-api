package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initializeRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/hello", hello(db))
	router.GET(fmt.Sprintf("/%v", os.Getenv("LOADER_IO_TOKEN")), func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("LOADER_IO_TOKEN"))
	})
	return router
}

func hello(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := db.Where("email_address in (?)", []string{"test@test.com"}).Find(&[]User{})
		c.JSON(http.StatusOK, query.Value)
	}
}
