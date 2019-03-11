package main

import "fmt"

func main()  {

LOOP:
	for i:=1; i<=3; i++ {
		for j:=1; j<=3; j++ {
			fmt.Printf("i = %d, j = %d \n", i, j)
			if j == 2 {
				// continue 外部循环
				continue LOOP
			}
		}
	}
}
/*
i = 1, j = 1
i = 1, j = 2
i = 2, j = 1
i = 2, j = 2
i = 3, j = 1
i = 3, j = 2
*/