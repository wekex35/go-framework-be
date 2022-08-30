package auth

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
	create(c.Request.Body)
}

func RetrieveAllUser(c *gin.Context) {
	findAll()
}

func RetrieveUser(c *gin.Context) {
	id := c.Param("id")
	findOne(id)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	update(id, c.Request.Body)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	remove(id)
}
