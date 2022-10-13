package object

import (
	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type Material interface {
	Scatter(r ray.Ray, rec HitRecord) (bool, r3.Vec, ray.Ray)
}

func Reflect(v, n r3.Vec) r3.Vec {
	return r3.Sub(v, r3.Scale(2*util.Vec3Dot(v, n), n))
}
