package main

import "fmt"

func main()  {
	myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Length: 4	Capacity: 5 (10-5)
	mySlice1 := myArray[5:9]
	fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice1), cap(mySlice1))
	// Length: 3	Capacity: 10 (10-0)
	mySlice2 := myArray[:3]
	fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice2), cap(mySlice2))
	// Length: 3	Capacity: 6	(10-4)
	mySlice3 := myArray[4:7]
	fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice3), cap(mySlice3))
}
