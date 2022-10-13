package object

import (
	"math"

	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type Sphere struct {
	Center r3.Vec
	Radius float64
}

func NewSphere(center r3.Vec, radius float64) *Sphere {
	return &Sphere{Center: center, Radius: radius}
}

func (s *Sphere) Hit(r *ray.Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r3.Sub(r.Origin, s.Center)
	a := util.Vec3Dot(r.Dir, r.Dir)
	halfB := util.Vec3Dot(oc, r.Dir)
	c := util.Vec3Dot(oc, oc) - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}

	sqrtDiscriminant := math.Sqrt(discriminant)

	// Tries to find a root of the polynomial which lies between tMin and tMax
	root := (-halfB - sqrtDiscriminant) / a
	if root < tMin || root > tMax {
		root = (-halfB + sqrtDiscriminant) / a
		if root < tMin || root > tMax {
			return false
		}
	}

	rec.T = root
	rec.Point = r.At(root)
	// (rec.Point - s.Center) / s.Radius
	rec.SetFaceNormal(r, r3.Scale(1/s.Radius, r3.Sub(rec.Point, s.Center)))

	return true
}
