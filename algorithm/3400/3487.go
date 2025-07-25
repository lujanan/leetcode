package algorith_3400

func maxSum(nums []int) int {
	var minNum = -101
	var ans int
	var numMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			if _, ok := numMap[nums[i]]; ok {
				continue
			}
			ans+= nums[i]
			numMap[nums[i]] = 0

		} else {
			minNum = max(minNum, nums[i])
		}
	}

	if len(numMap) > 0 {
		return ans
	}
	return minNum
}
