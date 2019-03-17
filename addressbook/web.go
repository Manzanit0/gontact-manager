package addressbook

import "github.com/gin-gonic/gin"

var db = new(Addressbook)

func DefaultRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/list", listContactsHandler)
	r.POST("/add", addContactHandler)

	return r
}

func listContactsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"contacts": db.Contacts,
	})
}

func addContactHandler(c *gin.Context) {
	var contact Contact
	c.BindJSON(&contact)

	db.Add(contact)

	c.JSON(200, gin.H{
		"name":  contact.Name,
		"email": contact.Email,
	})
}
