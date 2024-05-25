// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

type personDB interface {
	save(person Person) error
}

type LocalDatabase struct{}

type InMemoryDatabase struct {
	m map[int]Person
}

func (db *LocalDatabase) save(person Person) error {
	// Implementation for saving a person to the local database
	fmt.Printf("Saving person in Local DB: %+v\n", person)

	return nil
}

func (db *InMemoryDatabase) save(person Person) error {
	// Implementation for saving a person to the map in meory
	fmt.Printf("Saving person in memory DB: %+v\n", person)
	mapd := db.m
	if mapd == nil {
		mapd = make(map[int]Person)
	}
	mapd[person.ID] = person
	return nil
}

func main() {
	db := new(LocalDatabase)
	person := Person{
		ID:   101,
		Name: "John D",
	}
	db.save(person)

	memorydb := new(InMemoryDatabase)
	memorydb.save(person)
}
