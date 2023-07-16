package main

import (
	"errors"
	"log"
	"os"
)

type PhoneBookEntry struct {
	Name 	string
	Surname string
	Phone 	string
}

type Tree struct {
	Root *Node
}

type Node struct {
	Data  		map[string]*PhoneBookEntry
	Children 	map[uint8]*Node
}

func (t *Tree) Insert(contact *PhoneBookEntry) error {
	if t == nil {
		return errors.New("cannot insert a value into a nil tree")
	}

	t.Root.insertName(contact.Name, contact)
	t.Root.insertName(contact.Surname, contact)

	return nil
}

func (n *Node) insertName(name string, contact *PhoneBookEntry) {
	nameLength := len(name)
	currentNode := n
	for i := 0; i < nameLength; i++ {
		index := name[i]
		if currentNode.Children[index] == nil {
			currentNode.Children[index] = &Node{Children: map[uint8]*Node{}, Data: map[string]*PhoneBookEntry{}}
		}
		currentNode = currentNode.Children[index]
		currentNode.Data[contact.Phone] = contact
	}
}

func (t *Tree) find(contact string) map[string]*PhoneBookEntry {
	nameLength := len(contact)
	current := t.Root
	for i := 0; i < nameLength; i++ {
		index := contact[i]
		if current.Children[index] == nil {
			return nil
		}
		current = current.Children[index]
	}

	return current.Data
}


func main() {
	var contacts = []PhoneBookEntry {
		{
			Name:    "Vladimir",
			Surname: "Ivanov",
			Phone:   "087 999 999 99",
		},
		{
			Name: "John",
			Surname: "Doe",
			Phone: "555 123 456",
		},
		{
			Name:    "Smith",
			Surname: "Johnson",
			Phone:   "444 333 222",
		},
		{
			Name:    "Joan",
			Surname: "Wilson",
			Phone:   "545 656 767",
		},
		{
			Name:    "Vlad",
			Surname: "Dracula",
			Phone:   "666666666",
		},
		{
			Name:    "Sara",
			Surname: "Jay",
			Phone:   "123543345",
		},
	}

	var tree = Tree{Root: &Node{Children: map[uint8]*Node{}, Data: map[string]*PhoneBookEntry{}}}
	for i := 0; i < len(contacts); i++ {
		err := tree.Insert(&contacts[i])
		if err != nil {
			return
		}
	}

	if len(os.Args) > 1 {
		searchQuery := os.Args[1]

		results := tree.find(searchQuery)
		if results == nil {
			log.Println("No results found")
			return
		}

		for _, contact := range results {
			log.Printf("Name: %s, Surname: %s, Phone: %s", contact.Name, contact.Surname, contact.Phone)
		}
		return
	}

	log.Println("No search keyword provided")
}

