package object

import (
	"math"

	"github.com/gargakshit/plasticine/ray"
	"github.com/gargakshit/plasticine/util"
	"gonum.org/v1/gonum/spatial/r3"
)

type Dielectric struct {
	IndexOfRefraction float64
}

func NewDielectric(indexOfRefraction float64) *Dielectric {
	return &Dielectric{IndexOfRefraction: indexOfRefraction}
}

func (d *Dielectric) Scatter(r ray.Ray, rec HitRecord) (bool, r3.Vec, ray.Ray) {
	// TODO(AG): Fix dielectric materials
	attenuation := r3.Vec{X: 1, Y: 1, Z: 1}

	refractionRatio := d.IndexOfRefraction
	if rec.FrontFace {
		refractionRatio = 1 / refractionRatio
	}

	unitDir := r3.Unit(r.Dir)
	cosTheta := math.Min(util.Vec3Dot(r3.Scale(-1, unitDir), rec.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1

	if cannotRefract || reflectance(cosTheta, refractionRatio) > util.RealRand() {
		return true, attenuation,
			ray.NewRay(rec.Point, reflect(unitDir, rec.Normal))
	}

	return true, attenuation,
		ray.NewRay(rec.Point, refract(unitDir, rec.Normal, refractionRatio))
}

func refract(uv, n r3.Vec, etaIOverEtaN float64) r3.Vec {
	cosTheta := math.Min(util.Vec3Dot(r3.Scale(-1, uv), n), 1)
	rPerpendicular := r3.Scale(etaIOverEtaN, r3.Add(uv, r3.Scale(cosTheta, n)))
	rParallel := r3.Scale(-math.Sqrt(1-util.Vec3LengthSquared(rPerpendicular)), n)

	return r3.Add(rPerpendicular, rParallel)
}

func reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
