package main

import "fmt"

func main()  {
	salary := 55

	if salary >= 80 {
		fmt.Println("Happy!")
	} else if salary < 80 && salary >= 60 {
		fmt.Println("Fine...")
	} else {
		fmt.Println("So sad...")
	}
}

//So sad...