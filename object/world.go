package object

import (
	"gonum.org/v1/gonum/spatial/r3"
)

func CreateWorld() Hittable {
	return NewHitList([]Hittable{
		// Ground
		NewSphere(r3.Vec{Y: -100.5, Z: -1}, 100, NewLambertian(r3.Vec{X: 0.8, Y: 0.8})),
		// Left spheres
		NewSphere(r3.Vec{X: -1, Z: -1}, -0.45, NewDielectric(1.5)),
		NewSphere(r3.Vec{X: -1, Z: -1}, 0.5, NewDielectric(1.5)),
		// Middle sphere
		NewSphere(r3.Vec{Z: -1}, 0.5, NewLambertian(r3.Vec{X: 0.7, Y: 0.3, Z: 0.3})),
		// Right sphere
		NewSphere(r3.Vec{X: 1, Z: -1}, 0.5, NewMetal(r3.Vec{X: 0.8, Y: 0.8, Z: 0.8}, 0.8)),
	})
}
