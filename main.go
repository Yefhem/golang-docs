package main

import (
	"fmt"
	"strconv"
)

func main() {
	// i := strconv.FormatInt(-13, 10)
	// fmt.Printf("%v %T\n", i, i)

	// u := strconv.FormatUint(88, 10)
	// fmt.Printf("%v %T\n", u, u)

	b := strconv.FormatBool(true)
	fmt.Printf("%v %T\n", b, b)
}
