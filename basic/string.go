package main

import "fmt"

func main()  {

	var a string = "This is a \n"
	// \n 被原样输出
	var b string = `This is b \n`

	fmt.Printf("%s", a)

	fmt.Printf("%s", b)

	fmt.Printf("\n%T: %c", a[2], a[2])
}
/*
This is a
This is b \n
uint8: i
*/