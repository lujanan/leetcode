package algorithm_1600

import "testing"

func Test_maximalNetworkRank(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"t1", args{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}}, 4},
		{"t2", args{5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}}, 5},
		{"t3", args{8, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}}, 5},
		{"t4", args{15, [][]int{{8, 12}, {5, 11}, {5, 12}, {9, 4}, {0, 9}, {1, 8}, {10, 2}, {13, 14}, {3, 4}, {11, 3}, {11, 8}, {5, 10}}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maximalNetworkRank(tt.args.n, tt.args.roads); gotMax != tt.wantMax {
				t.Errorf("maximalNetworkRank() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
