package main

import "fmt"

func main() {

	fmt.Println("Exercise 1: Function Update")
	var pointerInt *int
	n := 44
	pointerInt = &n
	updateNwithPointer(pointerInt)
	fmt.Println("updated value is: ", n)

	fmt.Println("\nExercise 2: Structr pointer")

	john := user{"John", 20}
	johnPointerAddress := &john
	johnPointerAddress.updateAgeWithPointer()
	fmt.Println("updated Age is: ", *johnPointerAddress)

	fmt.Println("\nExercise 3: Method- value Receiver")
	u := user{"Mike", 50}
	u.updateName("Jack")
	fmt.Println(u.name)

	fmt.Println("\nExercise 4: Method -  pointer Receiver")
	u.updatenamewithPointer("Smith")
	fmt.Println(u.name)

}
func updateNwithPointer(n *int) {
	*n = *n + 1
}

type user struct {
	name string
	age  int
}

func (u *user) updateAgeWithPointer() {
	u.age = u.age + 1
}

func (u user) updateName(newName string) {
	u.name = newName
}
func (u *user) updatenamewithPointer(newName string) {
	u.name = newName
}
