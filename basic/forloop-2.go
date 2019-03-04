package main

import "fmt"

func main()  {
	a := 0
	b := 10

	// 类似while循环
	for a < b {
		a++
		fmt.Println("a=", a)
	}
}

/*
a= 1
a= 2
a= 3
a= 4
a= 5
a= 6
a= 7
a= 8
a= 9
a= 10
*/