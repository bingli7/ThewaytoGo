package main

import "fmt"

func main()  {
	a := 85
	switch {
	case a >= 80:
		fmt.Println("Great!")
		fallthrough
	case a >= 60 && a <80:
		fmt.Println("Just fine~")
	default:
		fmt.Println("Default!")
	}
}

/*
Great!
Just fine~
*/