package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Server Is Running")
}

func static(c *gin.Context) {
	c.String(http.StatusOK, "landing page will be available soon")
}

func mainRouter(router *gin.RouterGroup) {
	router.GET("/", static)
	router.GET("/health", healthCheck)
}

func initRouter() *gin.Engine {
	masterRouter := gin.Default()
	masterRouter.Use(cors.Default())
	routeGroup := masterRouter.Group("/")
	mainRouter(routeGroup)
	return masterRouter
}

func StartServer() {
	server := initRouter()
	server.Run(":8080")
}
