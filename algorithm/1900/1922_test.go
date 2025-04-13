// 1922.统计好数字的数目

package algorithm_1900

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{1}, 5},
		{"t2", args{4}, 400},
		{"t3", args{50}, 564908303},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodNumbers(tt.args.n); got != tt.want {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
