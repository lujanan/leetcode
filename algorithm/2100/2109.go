// 2109.向字符串添加空格

package algorithm_2100

func addSpaces(s string, spaces []int) string {
	var res = make([]byte, 0, len(s)+len(spaces))
	var pre int
	for _, idx := range spaces {
		if idx > pre {
			res = append(res, s[pre:idx]...)
		}
		res = append(res, ' ')
		pre = idx
	}

	res = append(res, s[spaces[len(spaces)-1]:]...)
	return string(res)
}
