package main

import "fmt"

func getReturn(a int) int {

	if a > 5 {
		return 1
	}

	return 2
}

func main()  {

	a := 1
	// return 2
	resp1 := getReturn(a)
	fmt.Println(resp1)

	b :=10
	// return 1
	resp2 := getReturn(b)
	fmt.Println(resp2)
}

/*
2
1
*/