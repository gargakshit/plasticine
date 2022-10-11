package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	width  = 256 * 8
	height = 256 * 8
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			img.Set(
				i, j,
				color.RGBA{
					R: uint8((float32(i) / float32(width)) * 256),
					G: uint8((float32(j) / float32(height)) * 256),
					B: 64,
					A: 255,
				},
			)
		}
	}

	f, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
