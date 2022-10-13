package object

import (
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type Lambertian struct {
	Albedo r3.Vec
}

func NewLambertian(albedo r3.Vec) *Lambertian {
	return &Lambertian{Albedo: albedo}
}

func (l Lambertian) Scatter(_ ray.Ray, rec HitRecord) (bool, r3.Vec, ray.Ray) {
	dir := r3.Add(rec.Normal, util.RandomUnitVec3())
	return true, l.Albedo, ray.NewRay(rec.Point, dir)
}
