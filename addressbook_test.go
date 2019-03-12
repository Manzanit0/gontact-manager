package gontact

import "testing"

func TestCreateContact(t *testing.T) {
	c := New("first last", "aa@bb.cc")

	if c.Name != "first last" {
		t.Errorf("The contact's first name doesn't match")
	}

	if c.Email != "aa@bb.cc" {
		t.Errorf("The contact's email doesn't match")
	}
}

func TestAddContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")

	addressbook.Add(contact)

	if addressbook.Contacts == nil {
		t.Errorf("The addressbook hasn't been initialized")
	}

	if len(addressbook.Contacts) <= 0 {
		t.Errorf("The addressbook doesn't have any contacts")
	}

	if addressbook.Contacts[0].Name != "name" {
		t.Errorf("The contact has the data wrong")
	}
}

func TestFindContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("name")

	if found == nil {
		t.Errorf("No contact found")
	}

	if found.Name != "name" {
		t.Errorf("Wrong contact found")
	}
}

func TestNotFoundContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("wrong name")

	if found != nil {
		t.Errorf("No contact should be found")
	}
}
