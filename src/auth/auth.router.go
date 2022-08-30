package auth

import "github.com/gin-gonic/gin"

func AuthRouter(router *gin.RouterGroup) {
	router.GET("/:id", RetrieveUser)
	router.GET("/", RetrieveAllUser)
	router.POST("/", CreateUser)
	router.PATCH("/", UpdateUser)
	router.DELETE("/", DeleteUser)
}
