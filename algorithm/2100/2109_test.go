// 2109.向字符串添加空格

package algorithm_2100

import "testing"

func Test_addSpaces(t *testing.T) {
	type args struct {
		s      string
		spaces []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"LeetcodeHelpsMeLearn", []int{8, 13, 15}}, "Leetcode Helps Me Learn"},
		{"t2", args{"icodeinpython", []int{1, 5, 7, 9}}, "i code in py thon"},
		{"t3", args{"spacing", []int{0, 1, 2, 3, 4, 5, 6}}, " s p a c i n g"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpaces(tt.args.s, tt.args.spaces); got != tt.want {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
