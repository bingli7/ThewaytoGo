package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	// 会有以下报错，因为a为array，而参数要求为slice
	// Cannot use 'a' (type [5]int) as type []int
	// sum := sliceAsArgv(a)
	// 因此传入slice：a[:]
	sum := sliceAsArgv(a[:])
	fmt.Println("sum =", sum)
}

func sliceAsArgv(a []int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}
/*
sum = 10
*/