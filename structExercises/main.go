package main

import "fmt"

type temperature float64
type car struct {
	car   string
	brand string
	year  string
}

type bankAccount struct {
	owner   string
	balance string
}

func main() {

	fmt.Println("\nExercise 1")
	type book struct {
		title  string
		author string
		price  float64
	}

	book1 := book{"The Twits", "Roahl Dahl", 4.99}
	fmt.Println(book1)

	fmt.Println("\nExercise 2")
	var temp1 temperature
	temp1 = 27
	fmt.Println(temp1)

	fmt.Println("\nExercise 3")
	var temp2 temperature = 45
	temp2.print()

	fmt.Println("\nExercise 4")
	c1 := car{brand: "Hyundai", year: "2025"}
	c1.PrintCarDetails()

	fmt.Println("\nExercise 5")

}

func (temp temperature) print() {
	fmt.Println("Temperature is : ", temp)
}

func (c car) PrintCarDetails() {
	fmt.Println("Brand: ", c.brand)
	fmt.Println("Year: ", c.year)
}

func (b bankAccount) deposit() {

}
