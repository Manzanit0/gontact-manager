package gontact

import "errors"

type Contact struct {
	Name  string
	Email string
}

type Addressbook struct {
	Contacts []Contact
}

func New(name, email string) Contact {
	return Contact{name, email}
}

func (addressBook *Addressbook) Add(contact Contact) (string, error) {
	if _, c := addressBook.findFirst(contact.Email); c != nil {
		return "", errors.New("gontact: a contact with the email " + contact.Email + " already exists")
	}

	addressBook.Contacts = append(addressBook.Contacts, contact)

	return contact.Email, nil
}

func (addressBook *Addressbook) Find(nameOrEmail string) []Contact {
	var foundContacts []Contact
	for i := 0; i < len(addressBook.Contacts); i++ {
		contact := addressBook.Contacts[i]
		if contact.Name == nameOrEmail || contact.Email == nameOrEmail {
			foundContacts = append(foundContacts, contact)
		}
	}

	return foundContacts
}

func (addrBook *Addressbook) Delete(email string) *Contact {
	index, contact := addrBook.findFirst(email)

	// Delete operation.
	if contact != nil {
		addrBook.Contacts = append(addrBook.Contacts[:index], addrBook.Contacts[index+1:]...)
	}

	return contact
}

func (addressBook *Addressbook) findFirst(email string) (int, *Contact) {
	var index int
	var contact *Contact

	contacts := addressBook.Contacts
	for i := 0; i < len(contacts); i++ {
		if contacts[i].Email == email {
			contact = &contacts[i]
			index = i
			break
		}
	}

	return index, contact
}
