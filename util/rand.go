package util

import (
	"github.com/valyala/fastrand"
)

const (
	uint32Max    uint32  = 1<<32 - 1
	uint32MaxF64 float64 = float64(uint32Max)
)

func RealRand() float64 {
	return float64(fastrand.Uint32()) / uint32MaxF64
}

func RealRandRange(min, max float64) float64 {
	return min + (max-min)*RealRand()
}
