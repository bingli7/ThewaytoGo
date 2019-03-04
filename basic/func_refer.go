package main

import "fmt"

func swap(a *int, b *int) {
	fmt.Println("Begin to swap:")
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println("End swap...")
}

func main() {
	i := 1111
	j := 3333
	fmt.Println("Before swap:","i=",i,",j=",j)
	swap(&i, &j)
	fmt.Println("After swap:","i=",i,",j=",j)
}

/*
Before swap: i= 1111 ,j= 3333
Begin to swap:
End swap...
After swap: i= 3333 ,j= 1111
*/