// 3479.水果成篮III

package algorith_3400

import "testing"

func Test_numOfUnplacedFruitsIII(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{4, 2, 5}, []int{3, 5, 4}}, 1},
		{"t2", args{[]int{3, 6, 1}, []int{6, 4, 7}}, 0},
		{"t3", args{[]int{4, 2, 5, 1, 3}, []int{3, 5, 4, 1, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruitsIII(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruitsIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
