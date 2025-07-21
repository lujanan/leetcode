// 1957.删除字符使字符串变好

package algorithm_1900

func makeFancyString(s string) string {
	var bin = []byte(s)
	var res = make([]byte, 0, len(s))
	var l byte
	var n = 0
	for i := 0; i < len(bin); i++ {
		if l != bin[i] {
			res = append(res, bin[i])
			l = bin[i]
			n = 1
		} else {
			if n == 2 {
				continue
			}
			n++
			res = append(res, bin[i])
		}
	}

	return string(res)
}
