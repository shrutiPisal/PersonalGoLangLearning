package main

import "fmt"

func main() {
	var intArray [3]int
	fmt.Println(intArray)

	var stringArray = [4]string{"a", "b", "c", "d"}
	fmt.Println(stringArray)

	var (
		floatArray [4]float64
		boolArray  [4]bool
	)
	fmt.Println(floatArray, boolArray)

	intArray2 := [5]int{2: 22, 4: 44}
	fmt.Println(intArray2)

	intArray3 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(intArray3)

	var products [1000000]int
	//printProducts(products)

	fmt.Println(len(products))
	fmt.Println(len(intArray))
	fmt.Println(len(stringArray))
	fmt.Println(len(floatArray))
	fmt.Println(len(boolArray))
	fmt.Println(len(intArray2))
	fmt.Println(len(intArray3))

	const a = 5
	newArray := [a]int{}
	fmt.Println(newArray)

	employeeNames := [...]string{"john", "peter"}
	fmt.Println("Names - ", employeeNames)
	fmt.Println("total employees: ", len(employeeNames))

}

func printProducts(products [1000000]int) {
	fmt.Println(products)
}
