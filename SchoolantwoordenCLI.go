package main

import (
	"fmt"
)

func main() {
	selectOpt()
}

func selectOpt() {
	fmt.Println("1. Get images from API \n2. Convert PDF to images ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	switch input {
	case "1":
		images, err := getImgFromApi()
		if err != nil {
			return
		}

		saveImages(images)
	case "2":
		images := PDFToImg()

		saveImages(images)

	default:
		fmt.Println("Wrong Input: " + input)
	}

}
