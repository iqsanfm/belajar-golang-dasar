package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai64)

	var name = "iqsan faisal"
	var i = name[0]
	var eString = string(i)

	fmt.Println(name)
	fmt.Println(i)
	fmt.Println(eString)
}