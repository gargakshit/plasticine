package vec

import (
	"math"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Float
}

type Vec3[T number] struct {
	Underlying [3]T
}

func NewVec3[T number](x T, y T, z T) *Vec3[T] {
	return &Vec3[T]{Underlying: [3]T{x, y, z}}
}

func (v *Vec3[T]) Add(num T) *Vec3[T] {
	v.Underlying[0] += num
	v.Underlying[1] += num
	v.Underlying[2] += num
	return v
}

func (v *Vec3[T]) Multiply(num T) *Vec3[T] {
	v.Underlying[0] *= num
	v.Underlying[1] *= num
	v.Underlying[2] *= num
	return v
}

func (v *Vec3[T]) Divide(num T) *Vec3[T] {
	return v.Multiply(1 / num)
}

func (v *Vec3[T]) At(i int) T {
	return v.Underlying[i]
}

func (v *Vec3[T]) AtPtr(i int) *T {
	return &v.Underlying[i]
}

func (v *Vec3[T]) LengthSquared() T {
	return v.At(0)*v.At(0) +
		v.At(1)*v.At(1) +
		v.At(2)*v.At(2)
}

func (v *Vec3[T]) Length() T {
	return T(math.Sqrt(float64(v.LengthSquared())))
}
