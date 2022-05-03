package main

import "fmt"

func main() {
	var arr1 = [...]string{"1", "2", "3"}
	var arr2 = [3]string{"1", "2", "3"}
	var arr3 = [3]string{"1", "2", "6"}
	// var arr4 = [3]int{1, 2, 3}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 == arr3)
	fmt.Println(arr2 == arr3)

	/* Böyle bir durumda hata alırız.
	 *	fmt.Println(arr3==arr4)
	 */
}
