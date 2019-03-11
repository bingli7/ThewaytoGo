package main

import "fmt"

func main() {
	result1, result2 := myFunc(2)
	fmt.Println(result1, result2)
}

func myFunc(input int) (x2 int, x3 int) {
	// 直接赋值
	x2 = 2 * input
	x3 = 3 * input
	// return后为空
	return
}
/*
4 6
*/