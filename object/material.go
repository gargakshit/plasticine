package object

import (
	"github.com/gargakshit/plasticine/ray"
	"gonum.org/v1/gonum/spatial/r3"
)

type Material interface {
	Scatter(r *ray.Ray, rec *HitRecord) (bool, r3.Vec, *ray.Ray)
}
