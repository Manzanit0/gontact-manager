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

func (addressBook *Addressbook) Find(name string) []Contact {
	var foundContacts = []Contact{}
	for i := 0; i < len(addressBook.Contacts); i++ {
		if addressBook.Contacts[i].Name == name {
			foundContacts =
				append(foundContacts, addressBook.Contacts[i])
		}
	}

	return foundContacts
}
