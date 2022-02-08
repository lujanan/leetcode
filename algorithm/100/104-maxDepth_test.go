package algorithm_100

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"t0",
			args{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			3,
		},
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
			2,
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
			2,
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
			3,
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
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 6,
								Left: &TreeNode{
									Val: 3,
								},
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			5,
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
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
