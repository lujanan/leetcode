package algorithm_600

import "testing"

func Test_findIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"t0", args{0}, 1},
		{"t1", args{1}, 2},
		{"t2", args{2}, 3},
		{"t3", args{3}, 3},
		{"t4", args{4}, 4},
		{"t5", args{5}, 5},
		{"t6", args{6}, 5},
		{"t7", args{7}, 5},
		{"t8", args{8}, 6},
		{"t9", args{9}, 7},
		{"t10", args{10}, 8},
		{"t11", args{11}, 8},
		{"t12", args{12}, 8},
		{"t13", args{13}, 8},
		{"t14", args{14}, 8},
		{"t15", args{15}, 8},
		{"t16", args{59034207}, 317811},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := findIntegers(tt.args.n); gotCount != tt.wantCount {
				t.Errorf("findIntegers() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
