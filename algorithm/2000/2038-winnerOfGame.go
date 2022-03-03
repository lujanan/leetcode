package algorithm_2000

// 2038. 如果相邻两个颜色均相同则删除当前颜色
// https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/

func winnerOfGame(colors string) bool {
	var cnt = [2]int{}
	for i, n := 0, len(colors); i < n; {
		var idx, char = i, colors[i]
		for i < n && colors[i] == char {
			i++
		}
		if i-idx > 2 {
			cnt[char-'A'] += i - idx -2
		}
	}
	return cnt[0] > cnt[1]
}

func winnerOfGame1(colors string) bool {
	var a, b int
	var t1 = colors[0]
	for i := 1; i < len(colors)-1; i++ {
		if t1 == colors[i] && t1 == colors[i+1] {
			if t1 == 'A' {
				a++
			} else {
				b++
			}
		} else {
			t1 = colors[i]
		}
	}
	return a > b
}
