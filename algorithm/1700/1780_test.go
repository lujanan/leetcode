// 1780.判断一个数字是否可以表示成三的幂的和

package algorithm_1700

import "testing"

func Test_checkPowersOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{12}, true},
		{"t2", args{91}, true},
		{"t3", args{21}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPowersOfThree(tt.args.n); got != tt.want {
				t.Errorf("checkPowersOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
