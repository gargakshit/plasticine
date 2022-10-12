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

// Lerp does a linear interpolation of two vectors using a provided factor. It
// is expected to operate on the factor range of [0.0, 1.0], but does not check
// for inclusion. May return unexpected results when the factor is out of the
// expected range
func Lerp(fac float64, v1, v2 r3.Vec) r3.Vec {
	return r3.Add(r3.Scale(1-fac, v1), r3.Scale(fac, v2))
}
