package main

import "fmt"

type blackList func(string) bool


func registerUser(name string, blackList blackList) {
	if blackList(name) {
		fmt.Println("You Are Blocked " + name)
	} else {
		fmt.Println("Welcome " + name)
	}
}

func main() {
	blackList := func (name string) bool {
		return name == "anjing"
	}
	registerUser("iqsan", blackList)
	registerUser("anjing", func (name string) bool  {
		return name == "anjing"
	})
}