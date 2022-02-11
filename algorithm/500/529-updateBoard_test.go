package algorithm_500

import (
	"reflect"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	type args struct {
		board [][]byte
		click []int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			"t1",
			args{
				[][]byte{
					{'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'M', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E'},
				},
				[]int{3, 0},
			},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			"t2",
			args{
				[][]byte{
					{'B', '1', 'E', '1', 'B'},
					{'B', '1', 'M', '1', 'B'},
					{'B', '1', '1', '1', 'B'},
					{'B', 'B', 'B', 'B', 'B'},
				},
				[]int{1, 2},
			},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.args.board, tt.args.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
