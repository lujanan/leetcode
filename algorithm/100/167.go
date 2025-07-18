// 167.两数之和II-输入有序数组

package algorithm_100

func twoSum(numbers []int, target int) []int {
	var l, r = 0, len(numbers) - 1
	for l < r {
		if numbers[l]+numbers[r] == target {
			break
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}

	return []int{l + 1, r + 1}
}
