package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 4}
	// Map的值为slice：[]int
	m1 := make(map[int][]int)
	m1[1] = s1
	fmt.Printf("%T\n", m1[1])
	for i, v := range m1[1] {
		fmt.Println(i, " ", v)
	}
	// Map的值为*slice：*[]int
	m2 := make(map[int]*[]int)
	m2[1] = &s1
	fmt.Printf("%T\n", m2[1])
	// 注意：range *m2
	for i, v := range *m2[1] {
		fmt.Println(i, " ", v)
	}
}
/*
[]int
0   0
1   1
2   2
3   3
4   4
*[]int
0   0
1   1
2   2
3   3
4   4
*/