package main

import "fmt"

func main()  {

	a := [5]int{0, 1, 2, 3, 4}
	b := a[:4]
	b[2] = 22
	// slice b引用自array a，若修改slice的值，相应array的值也会修改：因为slice是引用（指针）！
	for i,v := range a {
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