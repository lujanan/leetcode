// 1366. 通过投票堆团队排名

package algorithm_1300

import "testing"

func Test_rankTeams(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{[]string{"ABC", "ACB", "ABC", "ACB", "ACB"}}, "ACB"},
		{"t2", args{[]string{"WXYZ", "XYZW"}}, "XWYZ"},
		{"t3", args{[]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}}, "ZMNAGUEDSJYLBOPHRQICWFXTVK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankTeams(tt.args.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
