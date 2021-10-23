package algorithm_100

import "testing"

func Test_isSymmetric(t *testing.T) {
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
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 6,
							},
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 8,
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
						Right: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 6,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			},
			true,
		},
		{
			"t2",
			args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
