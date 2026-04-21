package main

import "log"

func main() {
	marksA, marksB := 90, 66
	if marksA >= 75 && marksB >= 75 {
		log.Println("First class")
	} else if marksA >= 50 || marksB >= 50 {
		log.Println("Second Class")
	} else if marksA >= 35 && marksB != 35 {
		log.Println("Third Class")
	} else {
		log.Println("Failed")
	}

	//average := marksA + marksB/2

	if average := marksA + marksB/2; average > 50 {
		log.Println("at least 50 marks")
	}
}
