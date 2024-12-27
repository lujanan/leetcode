// 80. 删除有序数组中的重复项 II

package algorithm_0

import "fmt"

func removeDuplicatesV2(nums []int) int {
	var idx, i, j int = 0, 0, 1
	for i < len(nums) && j < len(nums) {
		if j-i > 2 {
			fmt.Println(idx, i, j)
			for k := 0; k < 2 && i+k < j; k++ {
				if i != idx {
					nums[idx] = nums[i]
				}
				idx++
			}
			i = j
		}
		j++
		if nums[i] != nums[j] {
			i++
		}
	}

	fmt.Println(idx, i, j)
	if idx < i {
		for ; i < len(nums); idx, i = idx+1, i+1 {
			nums[idx] = nums[i]
		}
	}

	fmt.Println(nums[:len(nums)-num])
	return len(nums) - num
}
