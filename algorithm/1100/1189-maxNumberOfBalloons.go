package algorithm_1100

// https://leetcode-cn.com/problems/maximum-number-of-balloons/

func maxNumberOfBalloons(text string) int {
	var sMap = [5]int{}
	for i := 0; i < len(text); i++ {
		if text[i] == 'a' {
			sMap[0]++
		} else if text[i] == 'b' {
			sMap[1]++
		} else if text[i] == 'l' {
			sMap[2]++
		} else if text[i] == 'o' {
			sMap[3]++
		} else if text[i] == 'n' {
			sMap[4]++
		}
	}

	sMap[2] /= 2
	sMap[3] /= 2
	var cnt = sMap[0]
	for i := 1; i < len(sMap); i++ {
		if cnt > sMap[i] {
			cnt = sMap[i]
		}
	}
	return cnt
}
