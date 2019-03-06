package main

import "fmt"

func main()  {

	myChannel := make(chan int, 5)

	myChannel <- 1
	myChannel <- 2

	fmt.Println(<- myChannel)
	fmt.Println(<- myChannel)
}
