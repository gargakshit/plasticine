package main

import (
	"gonum.org/v1/gonum/spatial/r3"
)

type Ray struct {
	Origin r3.Vec
	Dir    r3.Vec
}

func NewRay(origin r3.Vec, dir r3.Vec) *Ray {
	return &Ray{Origin: origin, Dir: dir}
}

func (r *Ray) At(t float64) r3.Vec {
	return r3.Add(r.Origin, r3.Scale(t, r.Dir))
}
