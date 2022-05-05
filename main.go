package main

import "fmt"

func main() {
	var myMap = map[string]string{
		"a": "black",
		"b": "blue",
		"c": "orange",
		"d": "purple",
		"e": "red",
		"f": "yellow",
	}

	for i, v := range myMap {
		fmt.Println(i, v)
	}
}
