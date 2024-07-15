package main

import "fmt"


type HashName interface {
	GetName() string
}

func SayHello(value HashName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}
type Animal struct{
	Name string
}
func (person Person) GetName() string {
	return person.Name
}
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "iqsan"}
	animal := Animal{Name: "kucing"}
	SayHello(animal)
	SayHello(person)

	
}