package gontact

type Contact struct {
	name  string
	email string
}

type Addressbook struct {
	contacts []Contact
}

func New_Contact(name, email string) *Contact {
	c := new(Contact)
	c.name = name
	c.email = email
	return c
}

func Add_Contact(a *Addressbook, c *Contact) *Addressbook {
	if a.contacts == nil {
		a.contacts = make([]Contact, 0)
	}

	new_list := append(a.contacts, *c)
	a.contacts = new_list
	return a
}
