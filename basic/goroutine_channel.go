package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _,v := range a {
		sum += v
	}
	c <- sum
}

func main()  {

	array1 := []int{1,2,3,4,5}
	array2 := []int{6,7,8,9,10}

	myChannel := make(chan int)
	// 发送第一个array的sum到channel
	go sum(array1, myChannel)
	// 发送第二个array的sum到channel
	go sum(array2, myChannel)
	// 分别取channel中的值（2个sum）
	mySum1, mySum2 := <- myChannel, <- myChannel

	fmt.Println(mySum1, mySum2)
}
/*
40 15
*/