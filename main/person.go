package main

type Person struct {
	id string
	name string
}

func NewPerson(id string,name string) *Person {
	return &Person{id:id ,name:name}
}
