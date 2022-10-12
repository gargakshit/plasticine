package main

import (
	"image/color"

	"gonum.org/v1/gonum/spatial/r3"
)

// VecToRGBA scales the input vector by 255 and then returns a new RGBA color
func VecToRGBA(vec r3.Vec) color.RGBA {
	v := r3.Scale(255, vec)
	return color.RGBA{
		R: uint8(v.X),
		G: uint8(v.Y),
		B: uint8(v.Z),
		A: 255,
	}
}
