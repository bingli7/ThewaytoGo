package main

import "fmt"

func main()  {
	// new的第一种方式
	var s1 *[]int = new([]int)
	// new的第二种方式
	s2 := new([]int)
	// 指针类型，指向slice类型([]int)
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)
}
/*
*[]int
*[]int
*/