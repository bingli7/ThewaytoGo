package main

import "fmt"

func main()  {

	i := 0
	j := 10
	// 如果i=5，跳到下一个for循环
	LOOP: for i <= j {
		i++
		if i == 5 {
			goto LOOP
		}
		fmt.Println(i)
	}
}
/*
输出没有5：
1
2
3
4
6
7
8
9
10
11
*/