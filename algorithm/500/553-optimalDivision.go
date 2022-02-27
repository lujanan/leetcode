package algorithm_500

import (
	"fmt"
	"strconv"
)

// 553. 最优除法
// https://leetcode-cn.com/problems/optimal-division/

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	var res = fmt.Sprintf("%d/(%d", nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		res += "/" + strconv.Itoa(nums[i])
	}
	return res + ")"
}
