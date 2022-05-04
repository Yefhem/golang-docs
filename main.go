package main

import "fmt"

func main() {
	fruits := [3]string{"orange", "banana", "apple"}

	fruits2 := fruits

	fmt.Printf("Fruits : %v\n", fruits)
	fmt.Printf("Fruits2 : %v\n", fruits2)

	fruits[0] = "watermelon"

	fmt.Printf("Fruits : %v\n", fruits)
	fmt.Printf("Fruits2 : %v\n", fruits2)

	fmt.Println()

	numbers := [5]int{10, 20, 30, 40, 50}

	// Elemanlar referans olarak iletiliyor
	numbers2 := &numbers

	fmt.Printf("Fruits : %v\n", numbers)
	fmt.Printf("Fruits2 : %v\n", *numbers2)

	numbers[3] = 700

	fmt.Printf("Fruits : %v\n", numbers)
	fmt.Printf("Fruits2 : %v\n", *numbers2)
}
