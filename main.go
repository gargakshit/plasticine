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
	width   = 480 * 3
	height  = 270 * 3
	samples = 64
)

var world = object.CreateWorld()

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	numPartitions := runtime.GOMAXPROCS(0)

	log.Println("Width:", width)
	log.Println("Height:", height)
	log.Println("Parallel: true")
	log.Println("Partitions:", numPartitions)
	log.Println("Samples:", samples)
	log.Println("Light bounces:", maxDepth)

	partitionHeight := height / numPartitions
	wg.Add(numPartitions)

	cam := camera.NewCamera(samples, width, height, img)

	timeStart := time.Now()
	for i := 0; i < numPartitions; i++ {
		go performRayTracing(&wg, i*partitionHeight, (i+1)*partitionHeight, cam)
	}

	// Compute the last partition if there are still scanlines left
	if (numPartitions*1)*partitionHeight < height {
		wg.Add(1)
		go performRayTracing(&wg, (numPartitions*1)*partitionHeight, height, cam)
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

	s := &runtime.MemStats{}
	runtime.ReadMemStats(s)
	log.Println("Allocs:", s.Alloc)
	log.Println("NumGC:", s.NumGC)
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

const maxDepth = 48

func rayColor(r ray.Ray, depth int) r3.Vec {
	if depth > maxDepth {
		return r3.Vec{}
	}

	if h, rec := world.Hit(r, 0.0000000001, infinity); h {
		hit, attenuation, scatter := rec.Mat.Scatter(r, rec)
		if hit {
			return util.Vec3Mul(attenuation, rayColor(scatter, depth+1))
		}

		return r3.Vec{}
	}

	unitDir := r3.Unit(r.Dir)
	fac := 0.5 * (unitDir.Y + 1.0)

	return util.Lerp(
		fac,
		r3.Vec{X: 1.0, Y: 1.0, Z: 1.0},
		r3.Vec{X: 0.5, Y: 0.7, Z: 1.0},
	)
}
