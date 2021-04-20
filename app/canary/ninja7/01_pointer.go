package main

import "fmt"

func main() {
	// Assign the values for new variables
	i, j := "gg", "yy"

	// Assing variables via pointer
	p := &i

	// derefrencing the address to get the value
	fmt.Println(*p, j, i)
	fmt.Println(p)

	// change the value via pointer refrence
	*p = "changeme"

	fmt.Println(i)
}
