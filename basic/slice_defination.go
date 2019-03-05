package main

import "fmt"

func main()  {

	myArray := [8]int{1,2,3,4,5,6,7,8}

	// 第一种定义方式
	var s1 []int = []int {1, 2, 3, 4, 5}

	// 第二种定义方式
	s2 := []int {6, 7, 8, 9, 10}

	// 第三种定义方式：make([]T, length, capacity)
	var s3 = make([]int, 5, 10)
	s3 = myArray[:]

	// 第四种定义方式
	s4 := make([]int, 2, 5)
	s4 = myArray[:]

	fmt.Println("1st slice, ", "length:", len(s1), ", capacity:", cap(s1))
	for i:=0; i<len(s1)	; i++ {
		fmt.Println(s1[i])
	}

	fmt.Println("2nd slice:", "length:", len(s2), ", capacity:", cap(s2))
	for i:=0; i<len(s2); i++ {
		fmt.Println(s2[i])
	}

	fmt.Println("3rd slice:", "length:", len(s3), ", capacity:", cap(s3))
	for i:=0; i<len(s3); i++ {
		fmt.Println(s3[i])
	}

	fmt.Println("4th slice:", "length:", len(s4), ", capacity:", cap(s4))
	for i:=0; i<len(s4); i++ {
		fmt.Println(s4[i])
	}
}
/*
1st slice,  length: 5 , capacity: 5
1
2
3
4
5
2nd slice: length: 5 , capacity: 5
6
7
8
9
10
3rd slice: length: 8 , capacity: 8
1
2
3
4
5
6
7
8
4th slice: length: 8 , capacity: 8
1
2
3
4
5
6
7
8
*/
