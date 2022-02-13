package algorithm_400

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}}, 1},
		{"t2", args{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}}, 2},
		{"t3", args{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}}, 3},
		{"t4", args{"AACCTTGG", "AATTCCGG", []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
