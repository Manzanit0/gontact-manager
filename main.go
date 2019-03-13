package main

import "github.com/gin-gonic/gin"

var addressbook = new(Addressbook)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", listContacts)

	return r
}

func listContacts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": addressbook.Contacts,
	})
}
