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

func TestAddExistingContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")

	addressbook.Add(contact)
	email, error := addressbook.Add(contact)

	if email != "" {
		t.Errorf("An email as been saved")
	}

	if error.Error() != "gontact: a contact with the email email already exists" {
		t.Errorf("Wrong error message")
	}

	if len(addressbook.Contacts) > 1 {
		t.Errorf("The addressbook has one too many contacts")
	}
}

func TestFindContactByName(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("name")

	if len(found) != 1 {
		t.Errorf("Wrong amount of contacts found")
	}

	if found[0].Name != "name" {
		t.Errorf("Wrong contact found")
	}
}

func TestFindContactByEmail(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("email")

	if len(found) != 1 {
		t.Errorf("Wrong amount of contacts found")
	}

	if found[0].Name != "name" {
		t.Errorf("Wrong contact found")
	}
}

func TestNotFoundContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("wrong name")

	if len(found) > 0 {
		t.Errorf("A contact has been found")
	}
}

func TestDeleteExistingContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found, _ := addressbook.Delete("email")

	if found == nil {
		t.Errorf("No contact returned")
	}

	if len(addressbook.Contacts) > 0 {
		t.Errorf("The addressbook still has a contact")
	}
}

func TestDeleteNonExistingContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found, err := addressbook.Delete("wrong email")

	if found != nil {
		t.Errorf("Unexpected contact found: %+v", found)
	}

	if msg := err.Error(); msg != "No contact found by email: wrong email" {
		t.Errorf("Wrong error message = '%s'", msg)
	}

	if amount := len(addressbook.Contacts); amount != 1 {
		t.Errorf("Unexpected amount of contacts in addressbook = %d", amount)
	}
}
