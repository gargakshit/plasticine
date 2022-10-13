package object

import (
	"gonum.org/v1/gonum/spatial/r3"
)

func CreateWorld() Hittable {
	return NewHitList([]Hittable{
		NewSphere(r3.Vec{Y: -100.5, Z: -1}, 100, NewLambertian(r3.Vec{X: 1})),
		NewSphere(r3.Vec{Z: -1}, 0.5, NewLambertian(r3.Vec{Y: 1})),
	})
}
