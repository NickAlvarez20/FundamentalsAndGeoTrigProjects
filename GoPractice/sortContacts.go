package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name   string
	Email  string
	Age    int
	Groups []string // list of groups the contact belongs to
}

func sortByLengthGroups(contacts []Contact) {
	sort.Slice(contacts, func(i, j int) bool {
		return len(contacts[i].Groups) > len(contacts[j].Groups)
	})
}

func sortByAge_LengthGroups(contacts []Contact) {
	sort.Slice(contacts, func(i, j int) bool {
		return len(contacts[i].Groups) == len(contacts[j].Groups) && contacts[i].Age < contacts[j].Age
	})
}

func sortByName_LengthGroups_Age(contacts []Contact) {
	sort.Slice(contacts, func(i, j int) bool {
		return len(contacts[i].Groups) == len(contacts[j].Groups) && contacts[i].Age == contacts[j].Age && contacts[i].Name < contacts[j].Name
	})
}

func sortContacts(contacts []Contact) []Contact {
	sortByLengthGroups(contacts)
	sortByAge_LengthGroups(contacts)
	sortByName_LengthGroups_Age(contacts)
	return contacts
}

func main() {
	contacts := []Contact{}

	fmt.Println(sortContacts(contacts))
}
