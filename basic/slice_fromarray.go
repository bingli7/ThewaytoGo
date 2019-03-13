// 切片（slice）是对数组一个连续片段的引用，切片slice是引用类型

package main

import "fmt"

func main()  {

	myArray := [5]int{1, 2, 3, 4, 5}
	// 两种方式赋值，但mySlice2并不是slice
	mySlice1 := myArray[:]
	mySlice2 := myArray
	// 修改array的某个值
	myArray[2] = 22
	// mySlice1: []int切片类型
	fmt.Printf("mySlice1: %T\n", mySlice1)
	// myArray[:]: slice的值也会随着array改变
	// 注意：slice指向原来的array！！！
	for i,v := range mySlice1 {
		fmt.Println(i, " ", v)
	}
	// mySlice2: [5]int 跟myArray一样是array类型
	fmt.Printf("mySlice2: %T\n", mySlice2)
	// myArray: slice的值不会随着array改变
	for i,v := range mySlice2 {
		fmt.Println(i, " ", v)
	}
}
/*
mySlice1: []int
0   1
1   2
2   22
3   4
4   5
mySlice2: [5]int
0   1
1   2
2   3
3   4
4   5
*/