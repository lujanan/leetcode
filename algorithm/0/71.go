// 71. 简化路径

package algorithm_0

import "strings"

func simplifyPath(path string) string {
	var strArr []string
	var chars []byte
	var pushStr = func() {
		if len(chars) < 1 {
			return
		}

		if string(chars) == ".." {
			if len(strArr) > 0 {
				strArr = strArr[:len(strArr)-1]
			}
		} else if string(chars) != "." {
			strArr = append(strArr, string(chars))
		}
	}

	for _, v := range []byte(path) {
		if v == '/' {
			pushStr()
			chars = chars[:0]
			continue
		}
		chars = append(chars, v)
	}
	pushStr()

	return "/" + strings.Join(strArr, "/")
}
