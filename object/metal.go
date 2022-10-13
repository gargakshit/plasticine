package object

import (
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type Metal struct {
	Albedo r3.Vec
	Fuzz   float64
}

func NewMetal(albedo r3.Vec, fuzz float64) *Metal {
	return &Metal{Albedo: albedo, Fuzz: fuzz}
}

func (m *Metal) Scatter(r ray.Ray, rec HitRecord) (bool, r3.Vec, ray.Ray) {
	reflected := reflect(r3.Unit(r.Dir), rec.Normal)
	return util.Vec3Dot(reflected, rec.Normal) > 0,
		m.Albedo,
		ray.NewRay(rec.Point, r3.Add(reflected, r3.Scale(m.Fuzz, util.RandomUnitVec3())))
}
