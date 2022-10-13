package object

import (
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type HitRecord struct {
	Point     r3.Vec
	Normal    r3.Vec
	T         float64
	FrontFace bool
}

func NewHitRecord() *HitRecord {
	return &HitRecord{}
}

func (h *HitRecord) SetFaceNormal(r *ray.Ray, outwardNormal r3.Vec) {
	h.FrontFace = util.Vec3Dot(r.Dir, outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = r3.Scale(-1, outwardNormal)
	}
}

type Hittable interface {
	Hit(r *ray.Ray, tMin, tMax float64, rec *HitRecord) bool
}
