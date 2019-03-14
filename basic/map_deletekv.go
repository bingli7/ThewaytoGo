// 删除Map键值对：delete(map, key)

package main

import "fmt"

func main()  {
	m1 := make(map[int]string)
	m1[1] = "This is a string."
	// 删除键值对：delete(map, key)
	delete(m1, 1)
	if _, isPresent := m1[1]; isPresent {
		fmt.Printf("Yes, m1[1] exist.")
	} else {
		// 若map[key]已被删除：
		fmt.Printf("No, m1[1] does NOT exist.")
	}
}
/*
No, m1[1] does NOT exist.
*/