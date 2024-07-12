package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	var lulus2 bool = lulusNilaiAkhir || lulusAbsensi


	fmt.Println("hasil nilai =", lulus)
	fmt.Println("hasil nilai2 =", lulus2)
	
}