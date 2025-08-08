// 808.分汤

package algorithm_800

import "testing"

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"t1", args{50}, 0.62500},
		{"t2", args{100}, 0.71875},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); int64(got*100000) != int64(tt.want*100000) {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}
