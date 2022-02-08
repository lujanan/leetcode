package algorithm_0

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return nil
	}
	var numMap = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var res []string
	var fn func(idx int, sub string)
	fn = func(idx int, sub string) {
		if idx >= len(digits) {
			res = append(res, sub)
			return
		}
		if val, ok := numMap[digits[idx]]; ok {
			for _, c := range val {
				fn(idx+1, sub+string(c))
			}
		}
	}
	fn(0, "")
	return res
}
