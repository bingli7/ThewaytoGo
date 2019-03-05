package main

import (
	"errors"
	"fmt"
	"math"
)

func mySqrt(a float64) (float64, error) {

	if a == 0 {
		return 0, errors.New("Error! can't be sqrted by zero!")
	}
	return math.Sqrt(a), nil
}

func main() {

	f := 2.22
	result, err := mySqrt(f)
	fmt.Printf("The result of sqrt(%f) is: ", f)
	fmt.Println(result, err)

	f = 0
	result, err = mySqrt(f)
	fmt.Printf("The result of sqrt(%f) is: ", f)
	fmt.Println(result, err)
}
/*
The result of sqrt(2.220000) is: 1.489966442575134 <nil>
The result of sqrt(0.000000) is: 0 Error! can't be sqrted by zero!
*/