package main

import "fmt"

func main() {
	day := "tuesday"

	switch day {
	case "Monday":
		fmt.Println("calling task1")
	case "Tuesday", "tuesday":
		fmt.Println("calling task2")
	case "Wednesday":
		fmt.Println("calling task3")
	case "Thursday":
		fmt.Println("calling task4")
	case "Friday":
		fmt.Println("calling task5")
	case "Saturday":
		fmt.Println("calling task6")
	default:
		fmt.Println("calling default task")
	}

	marks := 22
	switch {
	case marks <= 30:
		fmt.Println("marks<=30 - Failed")
	case marks > 30 && marks < 50:
		fmt.Println("3rd class")
	case marks >= 50 && marks < 60:
		fmt.Println("2nd class")
	case marks >= 60:
		fmt.Println("1st class")
	}

	currency := "INR"
	switch currency {
	case "USD":
		fmt.Println("tasks for USD")
	case "EUR":
		fmt.Println("tasks for EUR")
		fallthrough
	case "GBP":
		fmt.Println("tasks for GBP")
	case "AUD":
		fmt.Println("tasks for AUD")
	default:
		fmt.Println("tasks for default currency")
	}
}
