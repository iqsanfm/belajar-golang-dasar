package main

import (
	"fmt"
)

func main() {

	name := "iqsan faisal"

	switch name {
	case "iqsan":
		fmt.Println("Hallo", name)
	case "ageng":
		fmt.Println("hallo", name)
	default:
		fmt.Println("Boleh berkenalan?", name)
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	name = "iqsan"
	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Anda Benar")

	}
}