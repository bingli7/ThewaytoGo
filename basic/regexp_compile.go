package main

import (
	"fmt"
	"regexp"
)

func main()  {
	pattern := "[0-5]*-[0-9]*"
	// Compile
	re, _ := regexp.Compile(pattern)
	myString := "This is my phone number: 010-6556865, you can keep it."
	// string中匹配到的字符串替换为"###-#####"
	newStr := re.ReplaceAllString(myString, "###-#####")
	fmt.Println(newStr)
}
/*
替换后：
This is my phone number: ###-#####, you can keep it.
*/