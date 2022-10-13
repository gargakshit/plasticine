package util

import (
	"testing"
)

func TestClamp(t *testing.T) {
	type args struct {
		num float64
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "num is less than min",
			args: args{
				num: 0,
				min: 10,
				max: 20,
			},
			want: 10,
		},
		{
			name: "num is more than max",
			args: args{
				num: 30,
				min: 10,
				max: 20,
			},
			want: 20,
		},
		{
			name: "num is between min and max",
			args: args{
				num: 15,
				min: 10,
				max: 20,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.num, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
