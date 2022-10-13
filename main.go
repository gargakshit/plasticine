package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gargakshit/plasticine/objects"
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
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

	var wg sync.WaitGroup
	numPartitions := runtime.GOMAXPROCS(0)
	log.Println("Running in parallel with", numPartitions, "partitions")

	partitionHeight := height / numPartitions
	wg.Add(numPartitions)

	world := objects.CreateWorld()

	timeStart := time.Now()
	for i := 0; i < numPartitions; i++ {
		go performRayTracing(img, &wg, i*partitionHeight, (i+1)*partitionHeight, world)
	}

	wg.Wait()
	log.Println("Ray tracing took", time.Since(timeStart))

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

func performRayTracing(
	img *image.RGBA, wg *sync.WaitGroup,
	startHeight, endHeight int,
	world objects.Hittable,
) {
	for j := startHeight; j < endHeight; j++ {
		for i := 0; i < width; i++ {
			u := float64(i) / (width - 1)
			v := float64(j) / (height - 1)
			r := ray.NewRay(
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

			color := rayColor(r, world)
			img.Set(i, height-j-1, util.VecToRGBA(color))
		}
	}

	wg.Done()
}

var infinity = math.Inf(1)

func rayColor(r *ray.Ray, world objects.Hittable) r3.Vec {
	hitRecord := objects.NewHitRecord()
	if world.Hit(r, 0, infinity, hitRecord) {
		// (normal + (1, 1, 1)) / 2
		return r3.Scale(0.5, r3.Add(hitRecord.Normal, r3.Vec{X: 1, Y: 1, Z: 1}))
	}

	unitDir := r3.Unit(r.Dir)
	fac := 0.5 * (unitDir.Y + 1.0)

	return util.Lerp(
		fac,
		r3.Vec{X: 1.0, Y: 1.0, Z: 1.0},
		r3.Vec{X: 0.5, Y: 0.7, Z: 1.0},
	)
}
