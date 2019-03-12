package main

import "fmt"

func main()  {
	// 传入函数名，回调函数调用double函数，并为其传入参数
	result := callback(5, double)
	fmt.Printf("%d\n", result)
}

// 函数作为参数：f func(a int) int
func callback(a int, f func(a int) int) int {
	return f(a)
}

func double(a int) int {
	return a * 2
}
/*
10
*/