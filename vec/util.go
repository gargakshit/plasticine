package vec

func Add[T number](a, b *Vec3[T]) *Vec3[T] {
	return NewVec3(
		a.Underlying[0]+b.Underlying[0],
		a.Underlying[1]+b.Underlying[1],
		a.Underlying[2]+b.Underlying[2],
	)
}

func Subtract[T number](a, b *Vec3[T]) *Vec3[T] {
	return NewVec3(
		a.Underlying[0]-b.Underlying[0],
		a.Underlying[1]-b.Underlying[1],
		a.Underlying[2]-b.Underlying[2],
	)
}

func Multiply[T number](a, b *Vec3[T]) *Vec3[T] {
	return NewVec3(
		a.Underlying[0]*b.Underlying[0],
		a.Underlying[1]*b.Underlying[1],
		a.Underlying[2]*b.Underlying[2],
	)
}

func MultiplyNum[T number](a *Vec3[T], b T) *Vec3[T] {
	return NewVec3(
		a.Underlying[0]*b,
		a.Underlying[1]*b,
		a.Underlying[2]*b,
	)
}

func DivideNum[T number](a *Vec3[T], b T) *Vec3[T] {
	return MultiplyNum(a, 1/b)
}

func DotProduct[T number](a, b *Vec3[T]) T {
	return a.Underlying[0]*b.Underlying[0] +
		a.Underlying[1]*b.Underlying[1] +
		a.Underlying[2]*b.Underlying[2]
}

func CrossProduct[T number](a, b *Vec3[T]) *Vec3[T] {
	return NewVec3(
		a.Underlying[1]*b.Underlying[2]-a.Underlying[2]*b.Underlying[1],
		a.Underlying[2]*b.Underlying[0]-a.Underlying[0]*b.Underlying[2],
		a.Underlying[0]*b.Underlying[1]-a.Underlying[1]*b.Underlying[0],
	)
}

func UnitVector[T number](a *Vec3[T]) *Vec3[T] {
	return DivideNum(a, a.Length())
}
