package algorithm_0

func canJumpV1(nums []int) bool {
	var  maxIdx, ln = 0, len(nums)
	for i := 0; i < ln; i++ {
		if i <= maxIdx {
			maxIdx = max(maxIdx, i+nums[i])
			if maxIdx >= ln - 1 {
				return true
			}
		}
	}
	return false
}

// 跳跃游戏
// https://leetcode-cn.com/problems/jump-game/
func canJump(nums []int) bool {
	var idx, maxIdx = 0, len(nums) - 1
	for i, v := range nums {
		if i > idx {
			// 无法到达当前位置
			return false
		} else if i+v >= maxIdx {
			// 可直接抵达终点
			return true
		}

		// 当前能到达最远的位置
		if idx < i+v {
			idx = i + v
		}
	}
	return false
}
