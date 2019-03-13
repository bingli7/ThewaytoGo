package main

import "fmt"

func main() {
	// 指定索引index为2和4的value
	a := [5]int{2: 22, 4: 44}
	for i, v := range a {
		fmt.Println(i, " ", v)
	}
}
/*
0   0
1   0
2   22
3   0
4   44
*/