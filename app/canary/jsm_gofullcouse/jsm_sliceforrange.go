package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	y := []string{"jsm", "royal", "singam"}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for j, v := range y {
		fmt.Println(j, v)
	}
}
