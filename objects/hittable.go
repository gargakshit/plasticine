package objects

import (
	"github.com/gargakshit/plasticine/ray"
	"gonum.org/v1/gonum/spatial/r3"
)

type HitRecord struct {
	Point  r3.Vec
	Normal r3.Vec
	T      float64
}

type Hittable interface {
	Hit(r *ray.Ray, tMin, tMax float64) bool
}
