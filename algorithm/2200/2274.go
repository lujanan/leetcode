//Alice 管理着一家公司，并租用大楼的部分楼层作为办公空间。Alice 决定将一些楼层作为 特殊楼层 ，仅用于放松。
//
// 给你两个整数 bottom 和 top ，表示 Alice 租用了从 bottom 到 top（含 bottom 和 top 在内）的所有楼层。另给你一个
//整数数组 special ，其中 special[i] 表示 Alice 指定用于放松的特殊楼层。
//
// 返回不含特殊楼层的 最大 连续楼层数。
//
//
//
// 示例 1：
//
//
//输入：bottom = 2, top = 9, special = [4,6]
//输出：3
//解释：下面列出的是不含特殊楼层的连续楼层范围：
//- (2, 3) ，楼层数为 2 。
//- (5, 5) ，楼层数为 1 。
//- (7, 9) ，楼层数为 3 。
//因此，返回最大连续楼层数 3 。
//
//
// 示例 2：
//
//
//输入：bottom = 6, top = 8, special = [7,6,8]
//输出：0
//解释：每层楼都被规划为特殊楼层，所以返回 0 。
//
//
//
//
// 提示
//
//
// 1 <= special.length <= 10⁵
// 1 <= bottom <= special[i] <= top <= 10⁹
// special 中的所有值 互不相同
//
//
// Related Topics 数组 排序 👍 27 👎 0

package algorithm_2200

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxConsecutive(bottom int, top int, special []int) int {
	// sort.Ints(special)
	// quickSort(special, 0, len(special)-1)
	mergeSort(special, 0, len(special)-1)

	var max = int(math.Max(float64(special[0]-bottom), float64(top-special[len(special)-1])))
	for i := 1; i < len(special); i++ {
		max = int(math.Max(float64(max), float64(special[i]-special[i-1]-1)))
	}
	return max
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := quickPartion(arr, l, r)
	quickSort(arr, l, mid-1)
	quickSort(arr, mid+1, r)
}

func quickPartion(arr []int, l, r int) int {
	var mid = arr[r]
	var i, j = l, r - 1
	for i <= j {
		if arr[i] <= mid {
			i++
			continue
		}
		if arr[j] >= mid {
			j--
			continue
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[r], arr[i] = arr[i], arr[r]
	return i
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)

	var tmpArr = append([]int{}, arr[l:r+1]...)
	for idx, i, j := l, 0, mid+1-l; idx <= r; idx++ {
		if i <= mid-l && (j > r-l || tmpArr[i] <= tmpArr[j]) {
			arr[idx] = tmpArr[i]
			i++
		} else {
			arr[idx] = tmpArr[j]
			j++
		}
	}
	fmt.Println(arr, l, r)
}

//leetcode submit region end(Prohibit modification and deletion)
