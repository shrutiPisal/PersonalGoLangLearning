package main

import "fmt"

func main() {
	fmt.Println("Exercise 1\n")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("\nExercise 2\n")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("\nExercise 3\n")
	fruits := []string{"apple", "banana", "orange"}
	for fruit := range fruits {
		fmt.Println(fruits[fruit])
	}

	fmt.Println("\nExercise 4\n")
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\nExercise 5\n")
	nums := []int{3, 6, 7, 10, 13, 14}
	n := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[n]%2 == 0 {
			count++
		}
		n++
	}
	fmt.Println("Even numbers:", count)
}
