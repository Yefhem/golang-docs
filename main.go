package main

import "fmt"

func main() {
	var x uint = 50 // 00110010
	var y uint = 30 // 00011110
	var z uint

	z = x & y
	fmt.Printf("x & y = %d binary = %b\n", z, z) // 00010010

	z = x | y
	fmt.Printf("x | y = %d binary = %b\n", z, z) // 00111110

	z = x ^ y
	fmt.Printf("x ^ y = %d binary = %b\n", z, z) // 00101100

	z = x << 1
	fmt.Printf("x << y = %d binary = %b\n", z, z) // 01100100

	z = x >> 1
	fmt.Printf("x >> y = %d binary = %b\n", z, z) // 00011001
}
