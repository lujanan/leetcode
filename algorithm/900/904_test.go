// 904.水果成篮

package algorithm_900

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 1}}, 3},
		{"t2", args{[]int{0, 1, 2, 2}}, 3},
		{"t3", args{[]int{1, 2, 3, 2, 2}}, 4},
		{"t4", args{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}}, 5},
		{"t5", args{[]int{6,2,1,1,3,6,6}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
