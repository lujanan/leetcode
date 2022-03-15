//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方
//式重构得到原数据。
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串
//反序列化为原始的树结构。
//
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的
//方法解决这个问题。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 10⁴] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 设计 字符串 二叉树 👍 784 👎 0

package algorithm_200

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var curr = []*TreeNode{root}
	var next []*TreeNode
	var strs []string
	for len(curr) > 0 {
		for _, node := range curr {
			if node == nil {
				strs = append(strs, "null")
			} else {
				strs = append(strs, strconv.Itoa(node.Val))
				next = append(next, node.Left, node.Right)
			}
		}
		curr, next = next, make([]*TreeNode, 0)
	}

	return strings.Join(strs, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var res *TreeNode
	strs := strings.Split(data, ",")
	if len(strs) > 0 {
		if strs[0] != "null" {
			num, err := strconv.Atoi(strs[0])
			if err != nil {
				return nil
			}
			res = &TreeNode{Val: num}
		} else {
			return nil
		}
	}

	var curr = []*TreeNode{res}
	var next []*TreeNode
	var i, sl = 1, len(strs)
	for len(curr) > 0 {
		for _, node := range curr {
			if i >= sl {
				break
			}

			if strs[i] != "null" {
				num, err := strconv.Atoi(strs[i])
				if err != nil {
					return nil
				}
				node.Left = &TreeNode{Val: num}
				next = append(next, node.Left)
			}

			if i+1 < sl && strs[i+1] != "null" {
				num, err := strconv.Atoi(strs[i+1])
				if err != nil {
					return nil
				}
				node.Right = &TreeNode{Val: num}
				next = append(next, node.Right)
			}
			i += 2
		}
		curr, next = next, make([]*TreeNode, 0)
	}

	return res
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
