package main

import "fmt"

func main()  {
	myString := []string{"Li", "Bing"}
	// defer 最后执行
	defer deferFunc(myString...)
	fmt.Printf("Program end...\n")
}

func deferFunc(str ...string)  {
	fmt.Printf("Deferring...")
	for _, value := range str {
		fmt.Printf("%s ", value)
	}
	fmt.Printf("\n")
}
/*
Program end...
Deferring...Li Bing
*/