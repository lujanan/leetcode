package algorithm_1000

import "testing"

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// {
		// 	"t1",
		// 	args{
		// 		[][]int{
		// 			{0, 0, 0, 0},
		// 			{1, 0, 1, 0},
		// 			{0, 1, 1, 0},
		// 			{0, 0, 0, 0},
		// 		},
		// 	},
		// 	3,
		// },
		// {
		// 	"t2",
		// 	args{
		// 		[][]int{
		// 			{0, 1, 1, 0},
		// 			{0, 0, 1, 0},
		// 			{0, 0, 1, 0},
		// 			{0, 0, 0, 0},
		// 		},
		// 	},
		// 	0,
		// },
		{
			"t3",
			args{
				[][]int{
					{0, 0, 0, 0},
					{1, 0, 1, 0},
					{0, 1, 1, 0},
					{0, 1, 0, 1},
					{1, 0, 0, 0},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numEnclaves(tt.args.grid); gotRes != tt.wantRes {
				t.Errorf("numEnclaves() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
