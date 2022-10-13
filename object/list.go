package object

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

func (h *HitList) Hit(r ray.Ray, tMin, tMax float64) (bool, HitRecord) {
	var rec HitRecord
	hit := false

	for _, obj := range h.Objects {
		if h, hr := obj.Hit(r, tMin, tMax); h {
			hit = true
			rec = hr
		}
	}

	return hit, rec
}
