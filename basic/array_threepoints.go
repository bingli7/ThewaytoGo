package main

import "fmt"

func main()  {
	a := [...]int{1, 2, 3, 4, 5}
	// 三个点[...]type的类型为： [5]int，而不是slice类型
	fmt.Printf("a: %T\n", a)
}