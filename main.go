package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"gonum.org/v1/gonum/spatial/r3"
)

const (
	width  = 1920
	height = 1080
	// Camera
	aspectRatio    = float64(width) / float64(height)
	viewportHeight = 2.0
	viewportWidth  = viewportHeight * aspectRatio
	focalLength    = 1.0
)

var (
	// Camera
	origin     = r3.Vec{}
	horizontal = r3.Vec{X: viewportWidth}
	vertical   = r3.Vec{Y: viewportHeight}
	focal      = r3.Vec{Z: focalLength}
	// origin - horizontal/2 - vertical/2 - vec(0, 0, focalLength)
	lowerLeftCorner = r3.Sub(
		r3.Sub(
			r3.Sub(origin, r3.Scale(0.5, horizontal)),
			r3.Scale(0.5, vertical),
		),
		focal,
	)
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			u := float64(i) / (width - 1)
			v := float64(j) / (height - 1)
			ray := NewRay(
				origin,
				// lowerLeftCorner + u*horizontal + v*vertical - origin
				r3.Sub(
					r3.Add(
						lowerLeftCorner,
						r3.Add(
							r3.Scale(u, horizontal),
							r3.Scale(v, vertical),
						),
					),
					origin,
				),
			)

			img.Set(i, height-j-1, VecToRGBA(rayColor(ray)))
		}
	}

	f, err := os.Create("out/out.png")
	if err != nil {
		log.Fatal(err)
	}
	// skipcq: GO-S2307
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Println("Error encoding the image to png:", err)
		return
	}
}

func rayColor(r *Ray) r3.Vec {
	unitDir := r3.Unit(r.Dir)
	fac := 0.5 * (unitDir.Y + 1.0)

	return Lerp(
		fac,
		r3.Vec{X: 1.0, Y: 1.0, Z: 1.0},
		r3.Vec{X: 0.5, Y: 0.7, Z: 1.0},
	)
}
