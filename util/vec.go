package util

import (
	"gonum.org/v1/gonum/spatial/r3"
)

// Vec3Dot computes the dot product of two vectors
func Vec3Dot(v1, v2 r3.Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func RandomVec3() r3.Vec {
	return r3.Vec{
		X: 0,
		Y: 0,
		Z: 0,
	}
}
