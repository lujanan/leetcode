package algorithm_400

import (
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]byte("a")}, 1},
		{"t2", args{[]byte("aabbccc")}, 6},
		{"t3", args{[]byte("abbbbbbbbbbbb")}, 4},
		{"t4", args{[]byte("aqwertyu")}, 8},
		{"t5", args{[]byte("aaaaaaaaaaaaaaa")}, 3},
		{"t6", args{[]byte("aa")}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
