package main

import (
	"bytes"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/ncruces/zenity"
	"image/jpeg"
	"strconv"
)

func PDFToImg() [][]byte {
	fmt.Println("Select PDF file:")
	filePath, err := zenity.SelectFile()
	if err != nil {
		return nil
	}

	fmt.Println("Selected: ", filePath)

	doc, err := fitz.New(filePath)
	if err != nil {
		return nil
	}

	var images [][]byte

	for n := 0; n < doc.NumPage(); n++ {

		if n == 0 {
			fmt.Print("Getting page " + strconv.Itoa(n))
		} else {
			fmt.Print("\rGetting page " + strconv.Itoa(n))
		}

		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		buf := new(bytes.Buffer)

		err = jpeg.Encode(buf, img, nil)
		if err != nil {
			return nil
		}

		images = append(images, buf.Bytes())
	}
	fmt.Print("\n")
	return images
}
