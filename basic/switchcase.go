package main

import "fmt"

func main()  {

	a := 80
	// 第一种方式：switch后接变量
	switch a {
	case 90:
		fmt.Println("Grade: A")
	case 80:
		fmt.Println("Grade: B")
	case 60:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Other Grade...")
	}

	score := 85
	// 第二种方式：switch后无变量，case后加判断：
	switch {
	case score >=80:
		fmt.Println("You are Great!")
	case score >=60 && score <80:
		fmt.Print("Just fine.")
	case score <60:
		fmt.Println("You need work harder~")
	default:
		fmt.Println("something else")
	}

	fmt.Println("End...")
}

/*
Grade: B
You are Great!
End...
*/