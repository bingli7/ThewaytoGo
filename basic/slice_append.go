package main

import "fmt"

func main()  {
	var a = []int {1,2,3,4,5}
	b := []int {6,7,8,9,10}

	a = append(a, 11,22)
	b = append(a, 777)
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
}

/*
[1 2 3 4 5 11 22]
[1 2 3 4 5 11 22 777]
*/