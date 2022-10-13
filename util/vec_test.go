package util

import (
	"testing"

	"gonum.org/v1/gonum/spatial/r3"
)

func TestVec3Dot(t *testing.T) {
	type args struct {
		v1 r3.Vec
		v2 r3.Vec
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "generates the correct dot product",
			args: args{
				v1: r3.Vec{X: 10, Y: 20, Z: 30},
				v2: r3.Vec{X: 40, Y: 50, Z: 60},
			},
			want: 3200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vec3Dot(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Vec3Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}
