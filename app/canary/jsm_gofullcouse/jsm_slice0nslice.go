package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x[0])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x) /// slicing a slice...
	fmt.Println(len(x))
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
