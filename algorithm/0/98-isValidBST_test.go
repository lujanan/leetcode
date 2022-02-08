package algorithm_0

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"t1",
			args{
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			true,
		},
		{
			"t2",
			args{
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			false,
		},
		{
			"t3",
			args{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			false,
		},
		{
			"t4",
			args{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			false,
		},
		{
			"t5",
			args{
				&TreeNode{
					Val: -2147483648,
					Right: &TreeNode{
						Val: 2147483647,
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
