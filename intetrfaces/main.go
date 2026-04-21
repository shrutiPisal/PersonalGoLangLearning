package main

func main() {
	/*intern1 := Intern{"John", 3333}
	intern2 := Intern{
		name:"Bob",
		id : 3333*/
}

/*var employee1, employee2 Employee
employee1 = intern1
employee2:= intern2*/
//}

type Employee interface {
	GetEmployeeID() int
}

type Intern struct {
	name string
	id   int
}
