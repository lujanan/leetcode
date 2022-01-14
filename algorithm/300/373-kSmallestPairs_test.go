//给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。
//
//
//
// 示例 1:
//
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
//
// 示例 2:
//
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
//
// 示例 3:
//
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//
//
// 提示:
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// -10⁹ <= nums1[i], nums2[i] <= 10⁹
// nums1 和 nums2 均为升序排列
// 1 <= k <= 1000
//
// Related Topics 数组 堆（优先队列） 👍 307 👎 0

package algorithm_300

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{[]int{1, 7, 11}, []int{2, 4, 6}, 3}, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{"t2", args{[]int{1, 1, 2}, []int{1, 2, 3}, 2}, [][]int{{1, 1}, {1, 1}}},
		{"t3", args{[]int{1, 2}, []int{3}, 3}, [][]int{{1, 3}, {2, 3}}},
		{"t4", args{[]int{1, 2, 7, 11, 12, 13, 14}, []int{2, 4, 5, 6, 7, 8, 9, 10}, 30}, [][]int{{1, 2}, {2, 2}, {1, 4}, {2, 4}, {1, 5}, {2, 5}, {1, 6}, {2, 6}, {1, 7}, {2, 7}, {1, 8}, {7, 2}, {1, 9}, {2, 8}, {1, 10}, {2, 9}, {7, 4}, {2, 10}, {7, 5}, {11, 2}, {7, 6}, {12, 2}, {7, 7}, {11, 4}, {7, 8}, {13, 2}, {7, 9}, {12, 4}, {14, 2}, {11, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
