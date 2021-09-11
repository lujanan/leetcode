//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
// 示例：
//
// 输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//
//
// 提示：
//
//
// 0 <= len(arr) <= 100000
// 0 <= k <= min(100000, len(arr))
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列）
// 👍 111 👎 0

package algorithm_0

import (
	"math/rand"
	"sort"
	"time"
)

func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	} else if len(arr) <= k {
		return arr
	}
	sort.Ints(arr)
	return arr[:k]

	rand.Seed(time.Now().UnixNano())
	var l, r = 0, len(arr) - 1
	for {
		m := parts(&arr, l, r)
		if (m == k) || (m+1) == k {
			return arr[:k]
		} else if m < k {
			l = m
		} else {
			r = m - 1
		}
	}
}

func parts(arr *[]int, l, r int) int {
	if l >= r {
		return l
	}
	mid := l + int(rand.Int63n(int64(r-l)))
	var mv = (*arr)[mid]
	(*arr)[mid] = (*arr)[r]
	var tmpl = l
	for j := l; j < r; j++ {
		if (*arr)[j] < mv {
			(*arr)[tmpl], (*arr)[j] = (*arr)[j], (*arr)[tmpl]
			tmpl++
		}
	}
	(*arr)[tmpl], (*arr)[r] = mv, (*arr)[tmpl]
	return tmpl
}

// buildHeap 数组建堆
func buildHeap(arr []int) []int {
	n := len(arr)
	if n > 0 {
		arr = append(arr, arr[0])
	}
	for i := n / 2; i >= 1; i-- {
		j := i
		maxPos := j
		for {
			if j*2 <= n && arr[maxPos] > arr[j*2] {
				maxPos = j * 2
			}
			if j*2+1 <= n && arr[maxPos] > arr[j*2+1] {
				maxPos = j*2 + 1
			}
			if maxPos == j {
				break
			}
			arr[j], arr[maxPos] = arr[maxPos], arr[j]
			j = maxPos
		}
	}
	return arr[1:]
}
