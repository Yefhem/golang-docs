package main

import "fmt"

var b string = "b"

func main() {
	var b string = "b2"

	if true {
		var b string = "b3"
		fmt.Println(b)
		if true {
			var b string = "b4"
			fmt.Println(b)
		}
	}

	fmt.Println(b)

	myLittleFunc()
}

func myLittleFunc() {
	fmt.Println(b)
}
