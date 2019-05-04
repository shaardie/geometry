package geometry

import "testing"

func Test_rayCastingAlgorithm(t *testing.T) {
	polygon := Polygon{
		Point{0, 0},
		Point{1, 0},
		Point{1, 1},
	}
	type args struct {
		polygon Polygon
		p       Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"in", args{polygon, Point{0.5, 0.25}}, 1},
		{"out", args{polygon, Point{-0.5, 0.5}}, -1},
		{"on", args{polygon, Point{0.5, 0.5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rayCastingAlgorithm(tt.args.polygon, tt.args.p); got != tt.want {
				t.Errorf("rayCastingAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
