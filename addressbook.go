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

func (addressBook *Addressbook) Add(contact Contact) {
	addressBook.Contacts = append(addressBook.Contacts, contact)
}

func (addressBook *Addressbook) Find(nameOrEmail string) []Contact {
	var foundContacts = []Contact{}
	for i := 0; i < len(addressBook.Contacts); i++ {
		contact := addressBook.Contacts[i]
		if contact.Name == nameOrEmail || contact.Email == nameOrEmail {
			foundContacts = append(foundContacts, contact)
		}
	}

	return foundContacts
}
