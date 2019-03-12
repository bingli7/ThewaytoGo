package main

import "fmt"

func main()  {
	pngSuffix := getAddSuffix("png")
	pdfSuffix := getAddSuffix("pdf")

	pngFileName := "picture"
	pdfFileName := "myCV"

	fmt.Println(pngSuffix(pngFileName))
	fmt.Println(pdfSuffix(pdfFileName))
}

func getAddSuffix(suffix string) func(filename string) string {
	return func(filename string) string {
		return filename + "." + suffix
	}
}
/*
picture.png
myCV.pdf
*/