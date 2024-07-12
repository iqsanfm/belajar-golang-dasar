package main

import "fmt"


func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b
	
	//Augment Assigments
var (
	i = 10
	o = 10
	p = 10
	q = 10
	r = 10
)
i += 10
o -= 10
p *= 10
q /= 10
r %= 10

fmt.Println("tambah", i)
fmt.Println("pengurangan",o)
fmt.Println("Perkalian",p)
fmt.Println("Pembagian",q)
fmt.Println("Sisa Pembagian",r)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//Unary Operator
	var j = 1
	j++
	fmt.Println("Unary Operator 1 =",j)
	j++
	fmt.Println("Unary Operator 2 =", j)
	j--
	fmt.Println("Unary Operator 3 pengurangan =", j)
	//


}