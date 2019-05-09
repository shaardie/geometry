package geometry

import (
	"reflect"
	"testing"
)

func Test_bentleyOttmannAlgorithm(t *testing.T) {
	type args struct {
		segments []Segment
	}
	tests := []struct {
		name       string
		args       args
		wantPoints []Point
		wantErr    bool
	}{
		{
			"simple",
			args{
				[]Segment{
					Segment{Point{10, 12}, Point{62, 93}},
					Segment{Point{12, 98}, Point{16, 85}},
					Segment{Point{91, 37}, Point{85, 36}},
				},
			},
			[]Point{Point{1, 2}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPoints, err := bentleyOttmannAlgorithm(tt.args.segments)
			if (err != nil) != tt.wantErr {
				t.Errorf("bentleyOttmannAlgorithm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPoints, tt.wantPoints) {
				t.Errorf("bentleyOttmannAlgorithm() = %v, want %v", gotPoints, tt.wantPoints)
			}
		})
	}
}
