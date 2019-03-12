package gontact

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

func AddContact(addressBook *Addressbook, contact Contact) {
	if addressBook.Contacts == nil {
		addressBook.Contacts = make([]Contact, 0)
	}

	addressBook.Contacts = append(addressBook.Contacts, contact)
}
