package algorithm_1400

import (
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", args{2}, []string{"1/2"}},
		{"t2", args{3}, []string{"1/2", "1/3", "2/3"}},
		{"t3", args{4}, []string{"1/2", "1/3", "1/4", "2/3", "3/4"}},
		{"t4", args{1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := simplifiedFractions(tt.args.n)
			res := make(map[string]bool)
			for _, v := range gotRes {
				res[v] = false
			}
			if len(res) != len(tt.wantRes) {
				t.Errorf("simplifiedFractions() = %v, want %v", gotRes, tt.wantRes)
			}
			for _, v := range tt.wantRes {
				if _, ok := res[v]; !ok {
					t.Errorf("simplifiedFractions() = %v, want %v", gotRes, tt.wantRes)
				}
			}

		})
	}
}
