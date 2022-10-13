package objects

import (
	"github.com/gargakshit/plasticine/ray"
)

type HitList struct {
	Objects []Hittable
}

func NewHitList(objects []Hittable) *HitList {
	return &HitList{Objects: objects}
}

func (h *HitList) Add(obj Hittable) {
	h.Objects = append(h.Objects, obj)
}

func (h *HitList) Clear() {
	h.Objects = h.Objects[:0]
}

func (h *HitList) Hit(r *ray.Ray, tMin, tMax float64, rec *HitRecord) bool {
	tmpRec := NewHitRecord()
	hit := false
	closest := tMax

	for _, obj := range h.Objects {
		if obj.Hit(r, tMin, tMax, tmpRec) {
			hit = true
			closest = tmpRec.T
			*rec = *tmpRec
		}
	}

	_ = closest

	return hit
}
