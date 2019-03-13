package main

import "fmt"

func main()  {
	// a的类型为 *[5]int
	a := new([5]int)
	a[2] = 2
	a[4] = 4
	// b的类型为[5]int
	b := [5]int{}
	// a赋值给b
	b = *a
	b[2] = 22
	b[4] = 44
	fmt.Printf("a's type: %T\n", a)
	// a和b可以用相同的range方法
	for index, value := range a {
		fmt.Println(index, " ", value)		// derefencing *a to get back to the array is not necessary!
	}
	fmt.Printf("b's type: %T\n", b)
	for index, value := range b {
		fmt.Println(index, " ", value)
	}
}
/*
a's type: *[5]int
0   0
1   0
2   2
3   0
4   4
b's type: [5]int
0   0
1   0
2   22
3   0
4   44
*/