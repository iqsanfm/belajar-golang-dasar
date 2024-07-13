package main

import "fmt"

func main() {

	name := "iqsanfaisal"

	if name == "iqsan" {
			fmt.Println("Hallo", name)
	} else if name == "ageng" {
		fmt.Println("Hello Ageng")
	} else {
		fmt.Println("Boleh berkenalan??", name)
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama anda sudah benar")
	}
}