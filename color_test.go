package main

import (
	"image/color"
	"reflect"
	"testing"

	"gonum.org/v1/gonum/spatial/r3"
)

func TestVecToRGBA(t *testing.T) {
	type args struct {
		vec r3.Vec
	}
	tests := []struct {
		name string
		args args
		want color.RGBA
	}{
		{
			name: "correctly converts a zero vector to RGBA",
			args: args{vec: r3.Vec{X: 0, Y: 0, Z: 0}},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		},
		{
			name: "correctly converts a vector to RGBA",
			args: args{vec: r3.Vec{X: 0.72, Y: 0.87, Z: 0.16}},
			want: color.RGBA{R: 183, G: 221, B: 40, A: 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VecToRGBA(tt.args.vec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VecToRGBA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerp(t *testing.T) {
	type args struct {
		fac float64
		v1  r3.Vec
		v2  r3.Vec
	}
	tests := []struct {
		name string
		args args
		want r3.Vec
	}{
		{
			name: "fac is 0",
			args: args{
				fac: 0,
				v1:  r3.Vec{X: 10, Y: 20, Z: 30},
				v2:  r3.Vec{X: 0, Y: 0, Z: 0},
			},
			want: r3.Vec{X: 10, Y: 20, Z: 30},
		},
		{
			name: "fac is 1",
			args: args{
				fac: 1,
				v1:  r3.Vec{X: 0, Y: 0, Z: 0},
				v2:  r3.Vec{X: 10, Y: 20, Z: 30},
			},
			want: r3.Vec{X: 10, Y: 20, Z: 30},
		},
		{
			name: "fac is a decimal",
			args: args{
				fac: 0.541,
				v1:  r3.Vec{X: 30, Y: 0, Z: 10},
				v2:  r3.Vec{X: 10, Y: 20, Z: 30},
			},
			want: r3.Vec{X: 19.18, Y: 10.82, Z: 20.82},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lerp(tt.args.fac, tt.args.v1, tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lerp() = %v, want %v", got, tt.want)
			}
		})
	}
}
