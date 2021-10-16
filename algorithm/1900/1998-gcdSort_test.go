package algorithm_1900

import "testing"

func Test_gcdSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{"t1", args{[]int{7, 21, 3}}, true},
		{"t2", args{[]int{5, 2, 6, 2}}, false},
		{"t3", args{[]int{10, 5, 9, 3, 15}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := gcdSort(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("gcdSort() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
