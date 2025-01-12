// 137. 只出现一次的数字 II

package algorithm_100

func singleNumber(nums []int) int {
	var res, tmp int
	for i := 0; i < 64; i++ {
		tmp = 0
		for j := 0; j < len(nums); j++ {
			if nums[j]&(1<<i) != 0 {
				tmp++
			}
		}
		if tmp%3 != 0 {
			res |= 1 << i
		}
	}
	return res
}
