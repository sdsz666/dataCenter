package controllers

import (
	"github.com/gin-gonic/gin"
	. "dataCenter/api/services"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//user
	{
		router.GET("/users", Users)

		router.POST("/user", Store)

		router.PUT("/user/:id", Update)

		router.DELETE("/user/:id", Destroy)

	}




	return router
}