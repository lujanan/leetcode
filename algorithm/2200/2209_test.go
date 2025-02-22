// 2209.用地毯覆盖后的最少白色砖块

package algorithm_2200

import "testing"

func Test_minimumWhiteTiles(t *testing.T) {
	type args struct {
		floor      string
		numCarpets int
		carpetLen  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"10110101", 2, 2}, 2},
		{"t2", args{"11111", 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumWhiteTiles(tt.args.floor, tt.args.numCarpets, tt.args.carpetLen); got != tt.want {
				t.Errorf("minimumWhiteTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
