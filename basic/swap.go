package main

import "fmt"

func swap(a, b *int)  {
	tmp := *a
	*a = *b
	*b = tmp
}

func main()  {

	a := 10
	b := 77

	fmt.Println("Before swap:")
	fmt.Println("a=", a, ", b=", b)

	swap(&a, &b)

	fmt.Println("After swap:")
	fmt.Println("a=", a, ", b=", b)
}

/*
Before swap:
a= 10 , b= 77
After swap:
a= 77 , b= 10
*/