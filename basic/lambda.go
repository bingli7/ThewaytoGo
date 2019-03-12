package main

import "fmt"

func main()  {
	// 定义匿名函数，并赋值给变量
	myLambda := func(a, b int) int {
		return a + b
	}
	// 通过变量名调用匿名函数
	sum := myLambda(10, 20)
	fmt.Printf("sum: %d \n", sum)
}
/*
sum: 30
*/