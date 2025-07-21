// 1957.删除字符使字符串变好

package algorithm_1900

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	var bin = []byte(s)
	var res = make([]byte, 0, len(s))
	res = append(res, bin[0], bin[1])
	var l1, l2 = bin[0], bin[1]
	for i := 2; i < len(bin); i++ {
		if bin[i] != l1 || bin[i] != l2 {
			res = append(res, bin[i])
			l1, l2 = l2, bin[i]
		}
	}

	return string(res)
}
