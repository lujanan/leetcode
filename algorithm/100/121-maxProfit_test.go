package algorithm_100

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"t2", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitDp(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"t2", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitDp(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
