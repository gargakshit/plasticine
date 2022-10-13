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

	"github.com/gargakshit/plasticine/camera"
	"github.com/gargakshit/plasticine/object"
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

const (
	width   = 1920
	height  = 1080
	samples = 128 * 2
)

var world = object.CreateWorld()

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	numPartitions := runtime.GOMAXPROCS(0)
	log.Println("Running in parallel with", numPartitions, "partitions with", samples, "samples")

	partitionHeight := height / numPartitions
	wg.Add(numPartitions)

	cam := camera.NewCamera(samples, width, height, img)

	timeStart := time.Now()
	for i := 0; i < numPartitions; i++ {
		go performRayTracing(&wg, i*partitionHeight, (i+1)*partitionHeight, cam)
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
	wg *sync.WaitGroup,
	startHeight, endHeight int,
	cam *camera.Camera,
) {
	for y := startHeight; y < endHeight; y++ {
		for x := 0; x < width; x++ {
			cam.Sample(x, y, rayColor)
		}
	}

	wg.Done()
}

var infinity = math.Inf(1)

func rayColor(r *ray.Ray) r3.Vec {
	hitRecord := object.NewHitRecord()
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
