// 2279. 装满石头的背包的最大数量

package algorithm_2200

import "testing"

func Test_maximumBags(t *testing.T) {
	type args struct {
		capacity        []int
		rocks           []int
		additionalRocks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2}, 3},
		{"t2", args{[]int{10, 2, 2}, []int{2, 2, 0}, 100}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBags(tt.args.capacity, tt.args.rocks, tt.args.additionalRocks); got != tt.want {
				t.Errorf("maximumBags() = %v, want %v", got, tt.want)
			}
		})
	}
}
