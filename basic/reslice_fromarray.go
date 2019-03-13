package main

import "fmt"

func main()  {
	myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySlice := myArray[6:]

	fmt.Printf("mySlice 6:9\n")
	for ix, value := range mySlice {
		fmt.Printf("%d, %d\n", ix, value)
	}
	// reslice重新分片
	// mySlice = mySlice[:6] 会报错，因为超过slice的capacity
	mySlice = mySlice[:4]
	fmt.Printf("After reslice to 0:3\n")
	for ix, value := range mySlice {
		fmt.Printf("%d, %d\n", ix, value)
	}
}
/*
mySlice 6:9
0, 6
1, 7
2, 8
3, 9
After reslice to 0:3
0, 6
1, 7
2, 8
3, 9
*/