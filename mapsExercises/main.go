package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Exercise 1: PhoneBook")
	phoneBook := make(map[string]string)
	phoneBook["John"] = "07654312345"
	phoneBook["Paul"] = "07654312356"
	phoneBook["George"] = "07654314567"

	for k, v := range phoneBook {
		fmt.Println("key is", k)
		fmt.Println("value is", v)
	}

	value, ok := phoneBook["Paul"]
	if ok {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Key not found")
	}

	delete(phoneBook, "George")
	fmt.Println(len(phoneBook))

	fmt.Println("\nExercise 2: Word Counter")
	input := "the cat sat on the mat the cat"

	words := strings.Fields(input)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	for word, count := range wordCount {
		fmt.Println("word:", word, "count:", count)
	}

	fmt.Println("\nExercise 3: NIL vs Empty Map")
	//snippet A
	a := map[string]int{}
	a["x"] = 10
	fmt.Println(a["x"])
	//Out put :10

	//snippet B
	var b map[string]int
	fmt.Println(b["x"]) // what happens here?: 0
	b["x"] = 10         // what happens here?: error assignment to entry in nil map

}
