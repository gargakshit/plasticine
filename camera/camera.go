package camera

import (
	"image"
	"math"

	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type RayColor func(ray.Ray, int) r3.Vec

type Camera struct {
	origin     r3.Vec
	lowerLeft  r3.Vec
	horizontal r3.Vec
	vertical   r3.Vec
	samples    int

	img    *image.RGBA
	width  int
	height int

	v          r3.Vec
	u          r3.Vec
	w          r3.Vec
	lensRadius float64
}

func NewCamera(
	samples, width, height int,
	vFov float64, lookFrom, lookAt, vUp r3.Vec,
	aperture, focusDist float64,
	img *image.RGBA,
) *Camera {
	theta := math.Pi * vFov / 180 // radians
	h := math.Tan(theta / 2)

	aspectRatio := float64(width) / float64(height)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := r3.Unit(r3.Sub(lookFrom, lookAt))
	u := r3.Unit(r3.Cross(vUp, w))
	v := r3.Cross(w, u)

	origin := lookFrom
	horizontal := r3.Scale(focusDist*viewportWidth, u)
	vertical := r3.Scale(focusDist*viewportHeight, v)
	// origin - horizontal/2 - vertical/2 - focusDist*w
	lowerLeft := r3.Sub(
		r3.Sub(
			r3.Sub(origin, r3.Scale(0.5, horizontal)),
			r3.Scale(0.5, vertical),
		),
		r3.Scale(focusDist, w),
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
		v:          v,
		u:          u,
		w:          w,
		lensRadius: aperture / 2,
	}
}

func (c *Camera) getRay(u, v float64) ray.Ray {
	rd := r3.Scale(c.lensRadius, util.RandomVec3Disk())
	offset := r3.Add(r3.Scale(rd.X, c.u), r3.Scale(rd.Y, c.v))

	dir := r3.Sub(
		r3.Sub(
			r3.Add(
				c.lowerLeft,
				r3.Add(
					r3.Scale(u, c.horizontal),
					r3.Scale(v, c.vertical),
				),
			),
			c.origin,
		),
		offset,
	)

	return ray.NewRay(r3.Add(c.origin, offset), dir)
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

func (c *Camera) Sample(x, y int, rayColor RayColor) {
	col := r3.Vec{}

	for i := 0; i < c.samples; i++ {
		u := (float64(x) + util.RealRand()) / float64(c.width)
		v := (float64(y) + util.RealRand()) / float64(c.height)

		col = r3.Add(col, rayColor(c.getRay(u, v), 0))
	}

	c.writePixel(x, y, col)
}
