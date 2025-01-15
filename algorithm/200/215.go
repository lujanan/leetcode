// 215. 数组中的第K个最大元素

package algorithm_200

// 桶排序
func findKthLargest(nums []int, k int) int {
	var bucket = make([]int, 20001)
	for _, v := range nums {
		bucket[v+10000]++
	}

	for i := 20000; i >= 0; i-- {
		k = k - bucket[i]
		if k <= 0 {
			return i - 10000
		}

	}
	return nums[0]
}

// 快排，超时了？？
func findKthLargestV2(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := partion(nums, l, r)
	quickSort(nums, l, mid-1)
	quickSort(nums, mid+1, r)
}

func partion(nums []int, l, r int) int {
	var num = nums[r]
	i, j := l, r-1
	for i <= j {
		if nums[i] >= num {
			i++
			continue
		}

		if nums[j] <= num {
			j--
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	nums[i], nums[r] = nums[r], nums[i]
	return i
}
