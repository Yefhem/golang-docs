package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

func main() {
	employee := employee{
		name:      "steve",
		age:       20,
		isMarried: true,
	}

	manager := manager{
		employee:  employee,
		hasDegree: true,
	}

	fmt.Println(employee)
	fmt.Println(manager)
	fmt.Println("s")
}
