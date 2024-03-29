package algorithm_600

import "testing"

func Test_shoppingOffers(t *testing.T) {
	type args struct {
		price   []int
		special [][]int
		needs   []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			"t1",
			args{
				[]int{2, 5},
				[][]int{{3, 0, 5}, {1, 2, 10}},
				[]int{3, 2},
			},
			14,
		},
		{
			"t2",
			args{
				[]int{2, 3, 4},
				[][]int{{1, 1, 0, 4}, {2, 2, 1, 9}},
				[]int{1, 2, 1},
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := shoppingOffers(tt.args.price, tt.args.special, tt.args.needs); gotRes != tt.wantRes {
				t.Errorf("shoppingOffers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
