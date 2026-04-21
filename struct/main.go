package main

import (
	"fmt"
)

type productID int64

func main() {
	type user struct {
		userName string
		age      int
		email    string
		salary   float64
	}

	var user1 user

	fmt.Println(user1)

	john := user{"John", 33, "john@abc.com", 4444444.33}
	fmt.Println(john)
	fmt.Printf("%V", john)

	bill := user{
		email:    "bill@abc.com",
		userName: "Bill",
		salary:   99.99,
	}
	fmt.Printf("%v", bill)

	fmt.Println("John's age is ", john.age)
	fmt.Println("bill's salary is ", bill.age)
	var peter user
	peter.userName = "Peter"
	peter.salary = 99.99
	peter.email = "peter@abc.com"
	fmt.Printf("%v", peter)

	type admin struct {
		employee   user
		department string
	}
	admin1 := admin{
		employee: user{
			userName: "john",
			age:      33,
		},
		department: "peopleOperations",
	}
	admin2 := admin{user{"peter", 33, "peter@abc.com", 99.00}, "HR"}

	fmt.Printf("%v", admin1)
	fmt.Printf("%v", admin2)

	fmt.Println("admin 1's name is ", admin1.employee.userName)

	var p productID
	p = 33333333
	productInfo(p)

	var v1 int64
	p = productID(v1)

	//alias
	type myInt = int64
	var v3 myInt = 123
	var v4 int64
	v3 = v4
	fmt.Printf("%d", v3)

}

func productInfo(pId productID) {

}
