package controllers

import (
	"github.com/gin-gonic/gin"
	. "myDemo/dataFactory/api/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	return router
}