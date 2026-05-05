package main

import "fmt"

func main() {
	fmt.Println("\nExercise 1: Multiply two Integers")
	result := multiplyInts(4, 5)
	fmt.Println("Multiplication of 4 and 5 is: ", result)

	fmt.Println("\nExercise 2: sum and difference")
	sum, minus := plusminus(5, 4)
	fmt.Printf("Sum: %d\nDifference: %d", sum, minus)

	fmt.Println("\nExercise 3: variadic function to find the maximum number")
	maxNum := maxNumber(1, 5, 3, 9, 2)
	fmt.Printf("Max number: %d\n", maxNum)

	fmt.Println("\nExercise 4: init function")
	fmt.Println("Program running")

}

func multiplyInts(x int, y int) int {
	return x * y
}

func plusminus(x int, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}

func maxNumber(values ...int) int {
	maximum := values[0]
	for i := 1; i < len(values); i++ {
		if maximum < values[i] {
			maximum = values[i]
		}
	}
	return maximum
}

func init() {
	fmt.Println("\nApplication starting")
}
