package main

import "fmt"

// func sayHelloto (firstName string, lastName string) {
// 	fmt.Println("Say Hello to", firstName, lastName)
// }
func sayHelloto (firstName string, lastName string, asalDari string, umurSaya int) {
	fmt.Println("Say Hello to", firstName, lastName, asalDari, umurSaya)
}
func main() {
	sayHelloto("iqsan", "Faisal", "garut", 24)
}