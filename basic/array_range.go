package main

import "fmt"

func main()  {
	// 三个点的用法（除可变参数外）：代表数组元素个数
	a := [...]string{"a", "b", "c", "d"}
	for i, v := range a {
		fmt.Println(i, " ", v)
	}
}