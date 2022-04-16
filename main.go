package main

import "fmt"

func init() {
	fmt.Println("init c")
}

func init() {
	fmt.Println("init a")
}

func init() {
	fmt.Println("init b")
}

func main() {
	fmt.Println("main")
}
