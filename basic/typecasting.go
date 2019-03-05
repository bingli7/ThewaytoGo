package main

import "fmt"

func main()  {
	var a int = 10
	var b int = 3

	c := float32(a)/float32(b)

	fmt.Printf("The result of 10/3 is: %f", c)
}

/*
The result of 10/3 is: 3.333333
*/