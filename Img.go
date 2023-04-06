package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func getImgFromApi() ([][]byte, error) {
	i := 1
	var images [][]byte

	fmt.Print("Enter the url \n---> ")
	var inputUrl string
	fmt.Scanln(&inputUrl)

	//make sure url ends with trailing slash
	if !strings.HasSuffix(inputUrl, "/") {
		inputUrl = inputUrl + "/"
	}

	for {
		// Build URL with current index
		url := inputUrl + strconv.Itoa(i) + ".jpg"
		if i == 1 {
			fmt.Print("Getting page " + strconv.Itoa(i))
		} else {
			fmt.Print("\rGetting page " + strconv.Itoa(i))
		}
		// Make HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			// Handle error
			fmt.Println("\nError:", err)
			return nil, err
		}
		image, err := io.ReadAll(resp.Body)
		if err != nil {
			// Handle error
			fmt.Println("\nError:", err)
			return nil, err
		}

		if resp.StatusCode == 404 {
			// Stop looping if 404 response is returned
			fmt.Println("\nFound a total of " + strconv.Itoa(i) + " images.")
			break
		} else {
			images = append(images, image)
			time.Sleep(50 * time.Millisecond)
			i++
		}

		// Close response body to free resources
		resp.Body.Close()
	}
	return images, nil
}

func saveImages(images [][]byte) {

	fmt.Print("Do you want to save the images? (y/n): \n---> ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	if strings.Contains(input, "y") || strings.Contains(input, "Y") {
		fmt.Println("Saving images...")
	} else if strings.Contains(input, "n") || strings.Contains(input, "N") {
		fmt.Println("Not saving images...")
		return
	}

	if _, err := os.Stat("images"); os.IsNotExist(err) {
		fmt.Println("Images directory does not exist, creating directory...")
		err := os.Mkdir("images", 0755)
		if err != nil {
			return
		}
	}

	fmt.Print("what alias do you want to give the images\n---> ")
	var alias string
	fmt.Scanln(&alias)

	var imgDir = "images/" + alias + "/"

	fmt.Println("Creating " + alias + " directory...")
	if _, err := os.Stat(imgDir); os.IsExist(err) {
		fmt.Println(alias + " already exists")
		return
	}
	err := os.Mkdir(imgDir, 0755)
	if err != nil {
		return
	}
	fmt.Println("Writing images to " + alias + " directory...")
	for i, image := range images {
		err := os.WriteFile(imgDir+strconv.Itoa(i+1)+".jpg", image, 0644)
		if err != nil {
			return
		}
	}

	fmt.Println("done")

}
