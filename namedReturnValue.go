package main

import "fmt"

func 	getCompleteName() (firstName, lastName, fullName string) {
	firstName = "iqsan"
	lastName = "faisal"
	fullName = "iqsanfaisal"
	return firstName, lastName, fullName
}

func main() {
	firstName, lastName, fullName := getCompleteName()
	fmt.Println(firstName, lastName, fullName)
}