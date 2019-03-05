package main

import "fmt"

// func参数、多个返回值，返回值只需指定类型（没有别名）
func swap(a, b string) (string, string) {
	return b, a
}

func printinfo(name string, age int) {
	fmt.Println(name, age)
}

func main()  {
	string1 := "First string"
	string2 := "Second string"
	s1, s2 := swap(string1, string2)
	fmt.Println("After swap, ", "1.", s1, "2.", s2)

	name := "libing"
	age :=18
	printinfo(name, age)
}

/*
After swap,  1. Second string 2. First string
libing 18
*/