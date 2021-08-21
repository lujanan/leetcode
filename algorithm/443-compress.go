package algorithm

import "strconv"

func compress(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}
	var idx, i, j = 0, 0, 1
	for j <= len(chars) {
		if j < len(chars) && chars[i] == chars[j] {
			j++
			continue
		}

		if chars[idx] != chars[i] {
			chars[idx] = chars[i]
		}
		idx++
		if j-i > 1 {
			numStr := []byte(strconv.Itoa(j - i))
			for n := 0; n < len(numStr); n++ {
				chars[idx] = numStr[n]
				idx++
			}
		}
		i = j
		j++
	}
	return idx
}
