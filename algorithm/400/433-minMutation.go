package algorithm_400

// https://leetcode-cn.com/problems/minimum-genetic-mutation/

// 单向搜索
// TODO 双向搜索
func minMutation(start string, end string, bank []string) int {
	var bMap = make(map[string]int)
	for _, v := range bank {
		bMap[v] = 1
	}

	var gl = []byte{'A', 'C', 'G', 'T'}
	var step int
	var curr = []string{start}
	var next []string
	var isCheck = make(map[string]int)

loop:
	for _, s := range curr {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(gl); j++ {
				var temp = []byte(s)
				if temp[i] == gl[j] {
					continue
				}
				temp[i] = gl[j]
				if string(temp) == end {
					return step + 1
				}

				if _, ok := isCheck[string(temp)]; ok {
					continue
				}
				if _, ok := bMap[string(temp)]; ok {
					next = append(next, string(temp))
					isCheck[string(temp)] = 1
				}
			}
		}
	}

	if len(next) > 0 {
		step++
		curr = append([]string{}, next...)
		next = next[:0]
		goto loop
	}
	return -1
}
