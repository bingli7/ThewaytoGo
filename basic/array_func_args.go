package main

import "fmt"

// 若形参为不定长array，则传入的array也需为不定长；定长array亦然
func getAverage(array []int, len int) float64  {
	sum := 0
	for i:=0; i<len; i++ {
		sum = sum + array[i]
	}
	return float64(sum)/float64(len)
}

func main() {
	myArray := []int{2, 4, 7, 9, 7, 3, 55, 34, 5, 87}
	avg := getAverage(myArray, 10)
	fmt.Println(avg)
}

// 21.3