package util

import (
	"math"

	"gonum.org/v1/gonum/spatial/r3"
)

// Vec3Dot computes the dot product of two vectors
func Vec3Dot(v1, v2 r3.Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Vec3LengthSquared(v r3.Vec) float64 {
	return Vec3Dot(v, v)
}

func Vec3Length(v r3.Vec) float64 {
	return math.Sqrt(Vec3LengthSquared(v))
}

func RandomVec3() r3.Vec {
	return r3.Vec{
		X: RealRand(),
		Y: RealRand(),
		Z: RealRand(),
	}
}

func RandomVec3Between(min, max float64) r3.Vec {
	return r3.Vec{
		X: RealRandRange(min, max),
		Y: RealRandRange(min, max),
		Z: RealRandRange(min, max),
	}
}

func RandomUnitVec3() r3.Vec {
	for {
		vec := RandomVec3Between(-1, 1)
		if Vec3LengthSquared(vec) < 1 {
			return vec
		}
	}
}
