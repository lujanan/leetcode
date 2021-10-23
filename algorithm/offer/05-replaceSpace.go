package algorithm_offer

func replaceSpace(s string) string {
	// return strings.ReplaceAll(s, " ", `%20`)
	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
