package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string{}
	// person["nama"] = "iqsan"
	// person["address"] = "garut"
	// fmt.Println(person)

	person := map[string]string{
		"name" : "iqsan",
		"address" : "garut",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "iqsan"
	book["ups"] = "Salah"

	fmt.Println("yang belum di hapus",book)

	delete(book, "ups")
	fmt.Println("yang sudah dihapus =", book)
	
}