package main

import "fmt"

func main() {
	// counter := 1

	// for counter <=10 {
	// 	fmt.Println("Perulangan Ke =", counter)
	// 	counter++
	// }
	// 	fmt.Println("Selesai")



	for 	counter := 1; counter <=10;	counter++{
		fmt.Println("Perulangan Ke =", counter)
	
	}
		fmt.Println("Selesai")

		names := []string{"iqsan", "faisal", "iqsanfaisal"}
		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
		}

		// for index, name := range names {
		// 	fmt.Println("index", index, "=", name)
		// }

		for _, name := range names {
			fmt.Println("index", "=", name)
		}
}