package main

import "fmt"

func main() {

	//make empty map
	employees := make(map[int]string)
	employees[1] = "John"
	employees[2] = "Paul"
	employees[3] = "George"

	for k, v := range employees {
		fmt.Println("key is", k)
		fmt.Println("value is", v)
	}

	id := 4
	fmt.Println("name of the id is", employees[id])

	//var NIL mao
	var salaries map[int]float64
	//salaries := map[int]float64{}
	salaries = map[int]float64{1: 333.00, 2: 444.00}
	fmt.Println(salaries)

	//map initialisation
	addresses := map[int]string{}
	addresses[1] = "street1"
	addresses[2] = "street2"
	fmt.Println(addresses)

	if addresses[1] == "London" {
		fmt.Println("address is london")
	} else {
		fmt.Println("address is not london")
	}

	rollcall := map[string]int{}
	rollcall["John"] = 1
	fmt.Println(rollcall)

}
