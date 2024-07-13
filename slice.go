package main

import (
	"fmt"
)

func main() {
	names := [...] string{"iqsan", "faisal", "iqsanfaisal", "iqsanfm", "gegeng", "ageng"}

	slice := names[4:6]
	slice1 := names[:3]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice1)

	//
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println("Belum dirubah =",days)

	daySlice1 := days[5:] //sabtu, minggu
	fmt.Println("daySlace1 =", daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("daySlaice Perubahan =", daySlice1)
	fmt.Println("Sudah diubah =", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2,5)
	newSlice[0] = "Iqsan"
	newSlice[1] = "faisal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Iqsanfaisal", "gegeng")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}