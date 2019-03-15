package addressbook

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
	email, err := addressbook.Add(contact)

	if email != "" {
		t.Errorf("An email as been saved = %s", email)
	}

	if msg := err.Error(); msg != "gontact: a contact with the email email already exists" {
		t.Errorf("Wrong error message = %s", msg)
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

	if amount := len(found); amount != 1 {
		t.Errorf("Wrong amount of contacts found = %d", amount)
	}

	if con := found[0]; con.Name != "name" && con.Email != "email" {
		t.Errorf("Wrong contact found = %+v", con)
	}
}

func TestFindContactByEmail(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("email")

	if amount := len(found); amount != 1 {
		t.Errorf("Wrong amount of contacts found = %d", amount)
	}

	if con := found[0]; con.Name != "name" && con.Email != "email" {
		t.Errorf("Wrong contact found = %+v", con)
	}
}

func TestNotFoundContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found := addressbook.Find("wrong name")

	if len(found) > 0 {
		t.Errorf("A contact has been found: %s", found)
	}
}

func TestDeleteExistingContact(t *testing.T) {
	addressbook := new(Addressbook)
	contact := New("name", "email")
	addressbook.Add(contact)

	found, err := addressbook.Delete("email")

	if err != nil {
		t.Errorf("Error returned: '%s'", err.Error())
	}

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
