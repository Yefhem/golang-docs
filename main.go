package main

import "fmt"

func main() {
	const a = 50
	const b = 50.5
	const c = "apple"
	const d = true

	var i int = 0

	fmt.Printf("%v - %T\n", a, a)
	fmt.Printf("%v - %T\n", b, b)
	fmt.Printf("%v - %T\n", c, c)
	fmt.Printf("%v - %T\n", d, d)

	i = a
	fmt.Println(i) 
}
