package main

import "fmt"

func main()  {
	a := 0
	// for后没有任何东西，类似：while(true)
	for {
		a ++
		fmt.Println("a = ", a)
		if a >=10 {
			break
		}
	}
}

/*
a =  1
a =  2
a =  3
a =  4
a =  5
a =  6
a =  7
a =  8
a =  9
a =  10
*/