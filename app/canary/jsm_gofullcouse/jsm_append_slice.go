package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	z := []int{1, 2, 3, 4, 5, 66, 77, 88}
	x = append(x, z...)
	fmt.Println(x)
}
