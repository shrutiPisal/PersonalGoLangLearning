package main

import "fmt"

func main() {

	//Exercise 1
	fmt.Println("Exercise1\n")
	var name string = "fName"
	var age int = 25
	var likeGo bool = true
	fmt.Printf("Name:%s\nAge:%d\nLike Go:%t\n", name, age, likeGo)

	//exercise 2
	fmt.Println("\nExercise2\n")
	var count int
	var price1 float64
	var username string
	var isAdmin bool
	fmt.Printf("Count:%d\nPrice:%f\nUsername:%s\nIsAdmin:%t\n", count, price1, username, isAdmin)

	//exercise 3
	fmt.Println("\nExercise3")
	const (
		companyName = "RI"
		maxRetries  = 3
		apiVersion  = "v1"
	)
	fmt.Printf("\nAPI Version:%s\nmaxRetries:%d\nCompany Name:%s\n ", apiVersion, maxRetries, companyName)

	//exercise 4
	//const age = 30
	//age =31 -> const value cannot be changed

	//exercise 5
	fmt.Println("\nExercise5")
	var productName string = "Pencil"
	var price2 float64 = 1.00
	var quantity float64 = 100
	var totalPrice float64 = price2 * quantity

	fmt.Printf("Total Price of %s : %f ", productName, totalPrice)

}
