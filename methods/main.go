package main

import "fmt"

type Employee struct {
	EMPLOYEEID int
	FirstName  string
}

func (e Employee) PrintEmployeeInfo() {
	fmt.Println("Employee ID: ", e.EMPLOYEEID)
	fmt.Println("Employee First Name: ", e.FirstName)
}

func (e Employee) GetEmployeeID() int {
	return e.EMPLOYEEID
}
func main() {
	e1 := Employee{EMPLOYEEID: 1, FirstName: "John"}
	e1.PrintEmployeeInfo()
	fmt.Println(e1.GetEmployeeID())

	var p productID = 4443
	p.PrintProductInfo()

	p2 := (productID)(5555)
	p2.PrintProductInfo()

}

type productID int64

func (p productID) PrintProductInfo() {
	fmt.Println("Product ID: ", p)
}
