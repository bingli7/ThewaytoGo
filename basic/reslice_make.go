// slice = slice[:length+1]
package main

import "fmt"

func main()  {
	a := make([]int, 5, 10)
	fmt.Printf("a's length: %d, capacity: %d\n", len(a), cap(a))

	// index out of range
	// fmt.Println(a[len(a)])

	// reslice: 扩展一位
	a = a[:len(a)+1]

	fmt.Printf("After slicing:\n")
	fmt.Printf("a's length: %d, capacity: %d\n", len(a), cap(a))
	// length超过capacity会报错
	// slice bounds out of range
	// a = a[:len(a)+5]
}