package main

import (
	"fmt"
	"regexp"
)

func main()  {
	pattern := "^[1-5]*-[6-9]*$"
	fmt.Println("Pattern:", pattern)

	myString1 := "123-768"
	ok, _ := regexp.Match(pattern, []byte(myString1))
	fmt.Println(myString1, ":", ok)

	myString2 := "123768"
	ok, _ = regexp.Match(pattern, []byte(myString2))
	fmt.Println(myString2, ":", ok)

	ok, _ = regexp.MatchString(pattern, myString1); fmt.Println(myString1, ":", ok)
	ok, _ = regexp.MatchString(pattern, myString2); fmt.Println(myString2, ":", ok)

}