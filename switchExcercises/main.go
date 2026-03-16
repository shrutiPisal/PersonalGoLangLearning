package main

import "fmt"

func main() {

	fmt.Println("Exercise 1")
	day := "Saturday"
	switch day {
	case "Monday":
		fmt.Println("Weekday")
	case "Tuesday":
		fmt.Println("Weekday")
	case "Wednesday":
		fmt.Println("Weekday")
	case "Thursday":
		fmt.Println("Weekday")
	case "Friday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	case "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("UnKnown")

	}

	fmt.Println("\nExercise 2")

	statusCode := 600
	switch statusCode {
	case 404:
		fmt.Println("404: Not Found")
	case 500:
		fmt.Println("500: Internal Server Error")
	case 200:
		fmt.Println("200: Success")
	case 301:
		fmt.Println("301: Moved Permanently")
	default:
		fmt.Println("UNKNOWN")
	}

	fmt.Println("\nExercise 3")

	Score := 59
	switch {
	case Score >= 90:
		fmt.Println("Grade A")
	case 89 >= Score && Score >= 80:
		fmt.Println("Grade B")
	case Score >= 70 && Score <= 79:
		fmt.Println("Grade C")
	case Score >= 60 && Score <= 69:
		fmt.Println("Grade D")
	case Score < 60:
		fmt.Println("Grade F")
	}

	fmt.Println("\nExercise 4")
	event := "EventCancelled"
	switch event {
	case "EventCancelled":
		fmt.Println("Event Cancelled")
		fallthrough
	case "EventReturnRequested":
		fmt.Println("Refund Initiated")
	case "EventDispatched":
		fmt.Println("Orders Shipped")
	default:
		fmt.Println("Unhandled Event")

	}
}
