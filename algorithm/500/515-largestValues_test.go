//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
//
//
//
// 示例1：
//
//
//
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//
//
// 示例2：
//
//
//输入: root = [1,2,3]
//输出: [1,3]
//
//
//
//
// 提示：
//
//
// 二叉树的节点个数的范围是 [0,10⁴]
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 168 👎 0

package algorithm_500

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			"t1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			[]int{1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := largestValues(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("largestValues() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
