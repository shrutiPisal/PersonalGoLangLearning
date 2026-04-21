package main

import "fmt"

func main() {
	/*

	 */
	var intvar1 int
	fmt.Println(intvar1)

	var intvar2 int8
	fmt.Println(intvar2)

	var intvar3 int16
	fmt.Println(intvar3)

	var intvar4 int32
	fmt.Println(intvar4)

	var intvar5 int64
	fmt.Println(intvar5)

	var intvar6 uint
	fmt.Println(intvar6)

	var intvar7 uint32
	fmt.Println(intvar7)

	var intvar8 uint64
	fmt.Println(intvar8)

	var intvar9 uint16
	fmt.Println(intvar9)

	var intvar10 uint8
	fmt.Println(intvar10)

	var age int
	age = 25
	fmt.Println(age)

	var age2 int = 32
	fmt.Println(age2)

	age3 := 44
	fmt.Println(age3)

	var salary float32
	var bonus float64
	fmt.Printf("salary: %.2f", salary)
	fmt.Printf("bonus: %.3f %T", bonus)

	tax := 35.00
	fmt.Printf("tax: %.4f type is %T", tax, tax)
}
