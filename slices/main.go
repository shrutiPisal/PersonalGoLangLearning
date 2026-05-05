package main

import "fmt"

func main() {

	var intSlice []int
	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	var floatSlice = []float64{33.2, 22.1}
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	stringSlice := []string{"first", "second", "third"}
	fmt.Println(stringSlice)
	fmt.Println(len(stringSlice))
	fmt.Println(cap(stringSlice))

	employeeIDs := make([]int, 10)
	fmt.Println("employeeID's ", employeeIDs)
	fmt.Println("total employee ID's", len(employeeIDs))
	fmt.Println("total capacity", cap(employeeIDs))

	employeeNames := make([]string, 10, 20)
	fmt.Println("employeeID's ", employeeNames)
	fmt.Println("total employee ID's", len(employeeNames))
	fmt.Println("total capacity", cap(employeeNames))

	intSlice = append(intSlice, 33)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	floatSlice = append(floatSlice, 13.2)
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	floatSlice = append(floatSlice, 44.3, 15.2)
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	newFloatSlice := floatSlice[2:4]
	fmt.Println(newFloatSlice)
	fmt.Println(len(newFloatSlice))
	fmt.Println(cap(newFloatSlice))

	newFloatSlice[1] = 99.9
	fmt.Println(newFloatSlice)
	fmt.Println(floatSlice)

	/* Original slice capacity is k
	newslice := oldslice[i:j]
	legnth of the new slice will be j-1
	capacity of the newslice is k-i
	*/

	arr := []int{1, 2, 3}
	arr = append(arr[:2], 4)
	fmt.Println(arr)

	numbers := []int{33, 44, 55}
	fmt.Println("Numbers: ", numbers)
	numbers2 := numbers[1:2]
	fmt.Println("Numbers2: ", numbers2)

	numbers2[0] = 22
	fmt.Println("updated numbers", numbers)
	fmt.Println("updated numbers2", numbers2)

}
