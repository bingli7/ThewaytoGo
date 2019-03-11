package main

import "fmt"

func main()  {
	firstname := "Bing"
	lastname := "Li"

	greeting(firstname, lastname)
}

func greeting(firstname, lastname string) int {

	fmt.Printf("%s %s\n", firstname, lastname)
	return 0
}
/*
Bing Li
*/