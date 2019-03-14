package main

import "fmt"

func main()  {
	m := make(map[string]string)
	if _, isPresent := m["name"]; isPresent {
		fmt.Printf("Yes, m[\"name\"] exists, and it's value is nil\n")
	} else {
		fmt.Printf("No, m[\"name\"] does NOT exist\n")
	}

}