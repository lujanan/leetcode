package algorithm_300

// 393. UTF-8 编码验证
// https://leetcode-cn.com/problems/utf-8-validation/

func validUtf8(data []int) bool {
	var n = len(data)
	for i := 0; i < n; {
		if data[i]>>7^1 == 1 {
			i++

		} else if data[i]>>5^6 == 0 {
			if i+1 >= n || data[i+1]>>6^2 != 0 {
				return false
			}
			i += 2

		} else if data[i]>>4^14 == 0 {
			if i+2 >= n {
				return false
			}
			for j := 1; j < 3; j++ {
				if data[i+j]>>6^2 != 0 {
					return false
				}
			}
			i += 3

		} else if data[i]>>3^30 == 0 {
			if i+3 >= n {
				return false
			}
			for j := 1; j < 4; j++ {
				if data[i+j]>>6^2 != 0 {
					return false
				}
			}
			i += 4

		} else {
			return false
		}
	}

	return true
}
