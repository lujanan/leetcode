package algorithm_1400

import "testing"

func Test_findMinFibonacciNumbers(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{7}, 2},
		{"t2", args{10}, 2},
		{"t3", args{19}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findMinFibonacciNumbers(tt.args.k); gotRes != tt.wantRes {
				t.Errorf("findMinFibonacciNumbers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
