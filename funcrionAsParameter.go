package main

import "fmt"

type filterString func(string) string


func sayHelloWithfilter(name string, filter filterString) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else if name == "anjing"{
		return "...."
	} else if name == "goblog" {
		return "....."
	} else{
		return name
	}

}


func main() {
		sayHelloWithfilter("iqsan", spamFilter)

		filter := spamFilter
		sayHelloWithfilter("goblog", filter)

}