package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name )
}

func main() {
	var iqsan Customer
	iqsan.Name = "Iqsan Faisal"
	iqsan.Address = "Indonesia"
	iqsan.Age = 24

	fmt.Println(iqsan)

	ageng := Customer{
		Name: "Ageng",
		Address: "Indonesia",
		Age: 24,
	}
	fmt.Println(ageng)

	faisal := Customer{"faisal", "indonesia", 24}
	fmt.Println(faisal)

	iqsan.sayHello("Budi")
	faisal.sayHello("agus")
	ageng.sayHello("ujang")

}