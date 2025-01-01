// 3280.将日期转换为二进制表示

package algorithm_3200

import "testing"

func Test_convertDateToBinary(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"2080-02-29"}, "100000100000-10-11101"},
		{"t2", args{"1900-01-01"}, "11101101100-1-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDateToBinary(tt.args.date); got != tt.want {
				t.Errorf("convertDateToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
