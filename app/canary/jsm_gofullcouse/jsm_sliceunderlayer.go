package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x, 43, 43, 43, 43, 43, 43, 44) // new underlying array allocated

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("Next Program2")
	two()
	fmt.Println("Next Program 3")
	three()
	fmt.Println("Next Program 4")
	four()
	


func two() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
	fmt.Println(y)
}



func three() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	_ = append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
}

func four() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
	fmt.Println("len", len(x))
	fmt.Println("cap", cap(x))

	fmt.Println(y)
	fmt.Println("len", len(y))
	fmt.Println("cap", cap(y))
}

