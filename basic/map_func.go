// func函数可以作为Map的值

package main

import "fmt"

func main() {
	// Map的value为func：func(i int) int
	funcMap := map[int]func(i int) int{
		1: func(i int) int { return i },
		10: func(i int) int { return i * 10 },
		100: func(i int) int { return i * 100 },
	}
	// 依次调用Map里的每个函数
	fmt.Println(funcMap[1](12))
	fmt.Println(funcMap[10](34))
	fmt.Println(funcMap[100](77))
}
/*
12
340
7700
*/