package main

import "github.com/gin-gonic/gin"
import "./addressbook"

var addrbook = new(addressbook.Addressbook)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/list", listContacts)
	r.POST("/add", addContact)

	return r
}

func listContacts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": addrbook.Contacts,
	})
}

func addContact(c *gin.Context) {
	var contact addressbook.Contact
	c.BindJSON(&contact)
	addrbook.Add(contact)
	c.JSON(200, gin.H{
		"name":  contact.Name,
		"email": contact.Email,
	})
}
