package main

import "fmt"

func main()  {
	const AAA = "my name is AAA"
	const (
		NAME = "libing"
		AGE = 18
	)

	fmt.Println(AAA)
	fmt.Println(NAME, AGE)
}

// my name is AAA
// libing 18