package algorithm

// 跳跃游戏
// https://leetcode-cn.com/problems/jump-game/
func canJump(nums []int) bool {
	var match int
	for i, v := range nums {
		if i > match {
			return false
		} else if match >= len(nums)-1 {
			return true
		}

		if match < i+v {
			match = i + v
		}
	}
	return false
}
