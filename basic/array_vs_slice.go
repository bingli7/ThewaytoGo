package main

import "fmt"

func main()  {

	a := [5]int{1, 2, 3, 4, 5}
	// type: [5]int
	fmt.Printf("a: %T\n", a)
	// type: []int
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("b: %T\n", b)
}
/*
a: [5]int
b: []int
*/