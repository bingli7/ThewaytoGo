package main

import "fmt"

func main()  {
	// 定义2个slice
	a := []int{0, 1, 2, 3, 4}
	b := a[1:3]
	fmt.Println("Printing b:")
	for i, v := range b {
		fmt.Println(i, " ", v)
	}
	a[2] = 22
	fmt.Println("Printing b after modifying a:")
	// slice b的值会改变，因为b指向a所指向的数组
	// 也可以这么理解：slice都是指针，a和b都指向同一块内存空间
	for i, v := range b {
		fmt.Println(i, " ", v)
	}
}
/*
Printing b:
0   1
1   2
Printing b after modifying a:
0   1
1   22
*/