//给定一个 N 叉树，找到其最大深度。
//
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
// N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：3
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：5
//
//
//
//
// 提示：
//
//
// 树的深度不会超过 1000 。
// 树的节点数目位于 [0, 10⁴] 之间。
//
// Related Topics 树 深度优先搜索 广度优先搜索 👍 237 👎 0

package algorithm_500

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name      string
		args      args
		wantDepth int
	}{
		{
			"t1",
			args{
				&Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 3,
							Children: []*Node{
								{
									Val:      5,
									Children: []*Node{},
								},
								{
									Val:      6,
									Children: []*Node{},
								},
							},
						},
						{
							Val:      2,
							Children: []*Node{},
						},
						{
							Val:      4,
							Children: []*Node{},
						},
					},
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDepth := maxDepth(tt.args.root); gotDepth != tt.wantDepth {
				t.Errorf("maxDepth() = %v, want %v", gotDepth, tt.wantDepth)
			}
		})
	}
}
