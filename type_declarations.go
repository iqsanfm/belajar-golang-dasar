package main

import "fmt"

func main() {
	type NoKTP string

	var ktpIqsan NoKTP = "1122334455"

	var contoh string = "222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpIqsan)
	fmt.Println(contohKtp)
}