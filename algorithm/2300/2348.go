// 2348. 全 0 子数组的数目

package algorithm_2300

func zeroFilledSubarray(nums []int) int64 {
	var ans, tmp int64
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			tmp++
		} else if tmp > 0 {
			ans += (tmp + 1) * tmp / 2
			tmp = 0
		}
	}
	if tmp > 0 {
		ans += (tmp + 1) * tmp / 2
	}

	return ans
}
