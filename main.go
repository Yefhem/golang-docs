package main

import (
	"fmt"
	"sort"
)

func main() {
	mySlice1 := []int{2, 5, 3, 7, 1, 9}
	mySlice2 := []string{"b", "a", "d", "c", "g", "e"}

	fmt.Println("Before Slice 1 :", mySlice1)
	fmt.Println("Before Slice 2 :", mySlice2)

	// sorting
	sort.Ints(mySlice1)
	sort.Strings(mySlice2)

	fmt.Println()

	fmt.Println("After Slice 1 :", mySlice1)
	fmt.Println("After Slice 2 :", mySlice2)
}
