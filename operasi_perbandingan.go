package main

import "fmt"

func main() {
	var name1 = "iqsan"
	var name2 = "iqsan"
	var name3 = "faisal"

	var result bool = name1 == name2
	var result2 bool = name1 == name3
	var result3 bool = name1 != name1
	var result4 bool = name1 != name3

	fmt.Println("sama dengan =", result)
	fmt.Println("sama dengan =", result2)
	fmt.Println("tidak sama dengan =", result3)
	fmt.Println("tidak sama dengan =",result4)
}