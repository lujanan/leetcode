package algorithm_1000

// 1010. 总持续时间可被 60 整除的歌曲
// https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/

func numPairsDivisibleBy60(time []int) int {
	var res int
	var list = make([]int, 60)
	for _, v := range time {
		var d = v % 60
		if d == 0 {
			res += list[d]
		} else {
			res += list[60-d]
		}
		list[v%60]++
	}
	return res
}
