package camera

import (
	"image"
	"math"

	"github.com/gargakshit/plasticine/object"
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

const (
	focalLength    = 1.0
	viewportHeight = 2.0
)

type RayColor func(*ray.Ray, *object.HitRecord, int) r3.Vec

type Camera struct {
	origin     r3.Vec
	lowerLeft  r3.Vec
	horizontal r3.Vec
	vertical   r3.Vec
	samples    int

	img    *image.RGBA
	width  int
	height int
}

func NewCamera(samples, width, height int, img *image.RGBA) *Camera {
	aspectRatio := float64(width) / float64(height)

	viewportWidth := aspectRatio * viewportHeight
	origin := r3.Vec{}
	horizontal := r3.Vec{X: viewportWidth}
	vertical := r3.Vec{Y: viewportHeight}
	// origin - horizontal/2 - vertical/2 - vec(0, 0, focalLength)
	lowerLeft := r3.Sub(
		r3.Sub(
			r3.Sub(origin, r3.Scale(0.5, horizontal)),
			r3.Scale(0.5, vertical),
		),
		r3.Vec{Z: focalLength},
	)

	return &Camera{
		origin:     origin,
		lowerLeft:  lowerLeft,
		horizontal: horizontal,
		vertical:   vertical,
		samples:    samples,
		img:        img,
		width:      width,
		height:     height,
	}
}

func (c *Camera) getRayDirection(u, v float64) r3.Vec {
	return r3.Sub(
		r3.Add(
			c.lowerLeft,
			r3.Add(
				r3.Scale(u, c.horizontal),
				r3.Scale(v, c.vertical),
			),
		),
		c.origin,
	)
}

func (c *Camera) writePixel(x, y int, color r3.Vec) {
	scale := 1.0 / float64(c.samples)
	color = r3.Vec{
		X: math.Sqrt(scale * color.X),
		Y: math.Sqrt(scale * color.Y),
		Z: math.Sqrt(scale * color.Z),
	}
	c.img.SetRGBA(x, c.height-y-1, util.VecToRGBA(color))
}

func (c *Camera) Sample(x, y int, r *ray.Ray, hr *object.HitRecord, rayColor RayColor) {
	col := r3.Vec{}

	for i := 0; i < c.samples; i++ {
		u := (float64(x) + util.RealRand()) / float64(c.width)
		v := (float64(y) + util.RealRand()) / float64(c.height)
		r.Origin = c.origin
		r.Dir = c.getRayDirection(u, v)

		col = r3.Add(col, rayColor(r, hr, 0))
	}

	c.writePixel(x, y, col)
}
