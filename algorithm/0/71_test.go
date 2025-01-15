// 71. 简化路径

package algorithm_0

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"/home/"}, "/home"},
		{"t2", args{"/home//foo/"}, "/home/foo"},
		{"t3", args{"/home/user/Documents/../Pictures"}, "/home/user/Pictures"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
