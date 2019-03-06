package main

import "fmt"

func main() {

	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5

	close(c)
	// 如果channel不关闭，则range不会结束，接收第6个数据时阻塞
	for i := range c {
		fmt.Println(i)
	}


}
