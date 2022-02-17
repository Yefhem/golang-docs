package main

import "fmt"

func main() {
	number := 101
	string1 := string(number)

	fmt.Printf("%v %T\n", number, number)
	fmt.Printf("%v %T\n", string1, string1)
}
