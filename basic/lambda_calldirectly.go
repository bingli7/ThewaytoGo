package main

import "fmt"

func main() {
	// (10, 20)为匿名函数的参数
	result := func(a, b int) int { return a + b }(10, 20)
	fmt.Printf("The result is：%d\n", result)
}
/*
The result is：30
*/