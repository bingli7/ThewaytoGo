package main

import "fmt"

func main()  {

	a := [5]int{0, 1, 2, 3, 4}
	// 修改array某个值
	a[2] = 22

	for i, v := range a {
		fmt.Println(i, " ", v)
	}
}
/*
0   0
1   1
2   22
3   3
4   4
*/