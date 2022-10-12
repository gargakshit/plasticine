package main

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/spatial/r3"
)

func TestNewRay(t *testing.T) {
	type args struct {
		origin r3.Vec
		dir    r3.Vec
	}
	tests := []struct {
		name string
		args args
		want *Ray
	}{
		{
			name: "generates a new ray",
			args: args{
				origin: r3.Vec{X: 10, Y: 20, Z: 30},
				dir:    r3.Vec{X: 5, Y: 10, Z: 1},
			},
			want: &Ray{
				Origin: r3.Vec{X: 10, Y: 20, Z: 30},
				Dir:    r3.Vec{X: 5, Y: 10, Z: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRay(tt.args.origin, tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRay_At(t *testing.T) {
	type fields struct {
		Origin r3.Vec
		Dir    r3.Vec
	}
	type args struct {
		t float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   r3.Vec
	}{
		{
			name: "generates the correct vector",
			fields: fields{
				Origin: r3.Vec{X: 10, Y: 20, Z: 30},
				Dir:    r3.Vec{X: 5, Y: 10, Z: 1},
			},
			args: args{5},
			want: r3.Vec{X: 35, Y: 70, Z: 35},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				Origin: tt.fields.Origin,
				Dir:    tt.fields.Dir,
			}
			if got := r.At(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}
		})
	}
}
