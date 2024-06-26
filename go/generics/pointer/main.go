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

type DBPointer[T any] interface {
	*T
	personDB
}

func CreatePersonDB[T any, dbPointer DBPointer[T]](person Person) error {

	var db dbPointer = new(T)
	db.save(person)
	return nil
}

func main() {
	person := Person{
		ID:   101,
		Name: "John D",
	}
	CreatePersonDB[LocalDatabase](person)
}
