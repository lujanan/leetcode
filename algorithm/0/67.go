// 67. 二进制求和

package algorithm_0

func addBinary(a string, b string) string {
	var res = make([]byte, max(len(a), len(b))+1)
	var ext, ai, bj uint8
	for i, j, k := len(a)-1, len(b)-1, len(res)-1; k >= 0; i, j, k = i-1, j-1, k-1 {
		ai, bj = 0, 0
		if i >= 0 {
			ai = a[i] - '0'
		}
		if j >= 0 {
			bj = b[j] - '0'
		}
		ext += ai + bj
		res[k] = byte(ext&1) + '0'
		ext >>= 1
	}

	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}
