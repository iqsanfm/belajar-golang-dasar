package main

import "fmt"


func getHello (name string) string{
	hello := "hello, nama saya " + name
	return hello
}


func main() {
	result := getHello("iqsan")

	fmt.Println(result)
	fmt.Println(getHello("iqsanfaisal"))
}