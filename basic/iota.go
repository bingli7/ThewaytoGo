package main

import "fmt"

func main()  {
	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C)
}

// 0 1 2