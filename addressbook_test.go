package gontact

import "testing"

func TestCreateContact(t *testing.T) {
	c := New_Contact("first last", "aa@bb.cc")

	if c.name != "first last" {
		t.Errorf("The contact's first name doesn't match")
	}

	if c.email != "aa@bb.cc" {
		t.Errorf("The contact's email doesn't match")
	}
}

func TestAddContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New_Contact("name", "email")

	a := Add_Contact(addressbook, contact)

	if a.contacts == nil {
		t.Errorf("The addressbook hasn't been initialized")
	}

	if len(a.contacts) <= 0 {
		t.Errorf("The addressbook doesn't have any contacts")
	}

	if a.contacts[0].name != "name" {
		t.Errorf("The contact has the data wrong")
	}
}
