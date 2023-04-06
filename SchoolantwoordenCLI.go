package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		selectOpt()
	}
}

func selectOpt() {
	fmt.Print("1. Get images from API \n2. Convert PDF to images\n3. Exit\n---> ")

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
	case "3":
		fmt.Println("Exiting")
		os.Exit(0)
	default:
		fmt.Println("Wrong Input: " + input)
	}

}
