package main

import "fmt"

func main() {
	var names [3] string
	names[0] = "iqsan"
	names[1] = "faisal"
	names[2] = "iqsan faisal"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int {
		9,
		10,
		11,
		12,
	}
	fmt.Println(values)
	fmt.Println(len(values))

}