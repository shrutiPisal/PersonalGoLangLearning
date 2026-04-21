package main

import (
	"fmt"

	"github.com/shrutiPisal/PersonalGoLandLearning/packages"

	"github.com/icrowley/fake"
)

func main() {
	result := addInts(2, 3)
	fmt.Println("result of addition of 2+3 is ", result)

	fmt.Printf("result of 4 +4 is %d", addInts(4, 4))

	/*Compilation error:
	var salary float32
	result = addInts(100, salary)
	var name string
	name = addInts(3,3)*/

	packages.CalculateDiscount(444)
	name, city, day, m, c := randomEmployeeDetails()
	fmt.Printf("Random Employee Details %s, %s, %d, %s, %s", name, city, day, m, c)

	printColour := func() {
		fmt.Println(fake.Color())
	}
	printColour()
}

func addInts(x int, y int) int {
	return x + y
}

func randomEmployeeDetails() (string, string, int, string, string) {
	name := fake.FullName()
	city := fake.City()
	day := fake.Day()
	month := fake.Month()
	colour := fake.Color()

	return name, city, day, month, colour
}
