package algorithm_300

// 393. UTF-8 编码验证
// https://leetcode-cn.com/problems/utf-8-validation/

// 优化
func validUtf8(data []int) bool {
	const (
		mask1 = 1 << 7
		mask2 = 1<<7 | 1<<6
		mask3 = 1<<7 | 1<<6 | 1<<5
		mask4 = 1<<7 | 1<<6 | 1<<5 | 1<<4
		mask5 = 1<<7 | 1<<6 | 1<<5 | 1<<4 | 1<<3
	)

	var n = len(data)
	for i := 0; i < n; {
		if data[i]&mask1 == 0 {
			i++

		} else if data[i]&mask3 == mask2 {
			if i+1 >= n || data[i+1]&mask1 != mask1 {
				return false
			}
			i += 2

		} else if data[i]&mask4 == mask3 {
			if i+2 >= n {
				return false
			}
			for j := 1; j < 3; j++ {
				if data[i+j]&mask2 != mask1 {
					return false
				}
			}
			i += 3

		} else if data[i]&mask5 == mask4 {
			if i+3 >= n {
				return false
			}
			for j := 1; j < 4; j++ {
				if data[i+j]&mask2 != mask1 {
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

func validUtf81(data []int) bool {
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
