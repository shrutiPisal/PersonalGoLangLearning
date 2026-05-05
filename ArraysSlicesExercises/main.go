package main

import "fmt"

func main() {
	fmt.Println("Exercise 1")
	var intArray = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Third Element: ", intArray[2])

	fmt.Println("\nExercise 2")
	intSlice := intArray[0:3]
	intSlice = append(intSlice, 40, 50)
	fmt.Println(intSlice)

	fmt.Println("\nExercise 3")

	intSlice2 := intArray[1:4]
	fmt.Println("SLice of the array is: ", intSlice2)

	fmt.Println("\nExercise 4")
	intSlice3 := make([]int, 2, 5)
	fmt.Println("Slice is: ", intSlice3)
	fmt.Println("Length is", len(intSlice3))
	fmt.Println("Capacity is", cap(intSlice3))

}
