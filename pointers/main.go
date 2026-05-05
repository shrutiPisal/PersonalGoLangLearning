package main

import "fmt"

func main() {
	var pointerInt *int
	age := 44
	fmt.Println("pointerInt", pointerInt)
	pointerInt = &age
	fmt.Println("pointerInt", pointerInt)

	fmt.Println("age value is: ", *pointerInt)

	age = 55
	fmt.Println("new age value is: ", *pointerInt)

	*pointerInt = 66
	fmt.Println("updated age value is: ", age)

	updateAgewithPointer(&age)
	fmt.Println("updated age value is: ", age)

	john := User{"John", 33}
	john.updateAge()
	fmt.Println("updated age for John is: ", john.age)

	johnAddress := &john
	johnAddress.updateAgeWithPointer()
	fmt.Println("updated age for john is: ", john.age)

}

func updateAgewithPointer(a *int) {
	*a = *a + 1
}

type User struct {
	name string
	age  int
}

func (u User) updateAge() {
	u.age = u.age + 1
}

func (u *User) updateAgeWithPointer() {
	u.age = u.age + 1
}
