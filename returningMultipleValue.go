package main

import "fmt"

func getFullName() (string, string) {
	return "iqsan", "faisal"
}
func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)


	//Mengambil 1 variable
	firstName, _ := getFullName()
	fmt.Println(firstName)
}