package main

import "fmt"

func main()  {
	// 三个点的用法（除可变参数外）：代表数组元素个数
	a := [...]string{"a", "b", "c", "d"}
	// 返回2个值，分别为index和value
	for i, v := range a {
		fmt.Println(i, " ", v)
	}
	// 也可只获得一个值（index）
	for i := range a {
		fmt.Println(i, " ", a[i])
	}
}
/*
0   a
1   b
2   c
3   d
0   a
1   b
2   c
3   d
*/