package main

import (
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	var format string
	flag.StringVar(&format, "f", "", "output format. support jpeg, png, gif")
	flag.Parse()
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	switch format {
	case "jpeg":
		err = jpeg.Encode(os.Stdout, img, nil)
	case "png":
		err = png.Encode(os.Stdout, img)
	case "gif":
		err = gif.Encode(os.Stdout, img, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}
