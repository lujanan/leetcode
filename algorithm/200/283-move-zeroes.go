//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 1175 👎 0

package algorithm_200

func moveZeroes(nums []int) []int {
	var length = len(nums)
	var i, j int
	for i < length {
		if nums[i] == 0 {
			if j <= i {
				j = i + 1
			}
			for ; j < length; j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					j++
					break
				}
			}
		}
		i++
	}
	return nums
}

func moveZeroesV2(nums []int) []int {
	var length = len(nums)
	if length < 2 {
		return nums
	}
	var i, j = 0, 1
	for {
		if j >= length {
			break
		}
		if nums[i] != 0 {
			i++
		} else if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return nums
}
