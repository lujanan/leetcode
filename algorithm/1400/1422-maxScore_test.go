package algorithm_1400

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"t1", args{"011101"}, 5},
		{"t2", args{"00111"}, 5},
		{"t3", args{"1111"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxScore(tt.args.s); gotMax != tt.wantMax {
				t.Errorf("maxScore() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
