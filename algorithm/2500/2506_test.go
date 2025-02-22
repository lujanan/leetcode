// 2506.统计相似字符串对的数目

package algorithm_2500

import "testing"

func Test_similarPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]string{"aba", "aabb", "abcd", "bac", "aabc"}}, 2},
		{"t2", args{[]string{"aabb", "ab", "ba"}}, 3},
		{"t3", args{[]string{"nba", "cba", "dba"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarPairs(tt.args.words); got != tt.want {
				t.Errorf("similarPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
