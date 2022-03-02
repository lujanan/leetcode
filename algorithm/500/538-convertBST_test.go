package algorithm_500

import (
	"reflect"
	"testing"
)

func Test_convertBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"t1",
			args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 7,
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			&TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 36,
					Left: &TreeNode{
						Val: 36,
					},
					Right: &TreeNode{
						Val: 35,
						Right: &TreeNode{
							Val: 33,
						},
					},
				},
				Right: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 26,
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
