// 1366. 通过投票堆团队排名

package algorithm_1300

import (
	"sort"
)

func rankTeams(votes []string) string {
	var arr = make([][27]int, 26)
	for _, ts := range votes {
		for i, ch := range ts {
			arr[ch-'A'][26] = int(ch-'A') + 1
			arr[ch-'A'][i]++
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		for k := 0; k < 26; k++ {
			if arr[i][k] == arr[j][k] {
				continue
			}
			return arr[i][k] > arr[j][k]
		}
		return arr[i][26] < arr[j][26]
	})

	var res []byte
	for i := 0; i < len(arr); i++ {
		if arr[i][26] <= 0 {
			break
		}
		res = append(res, byte(arr[i][26]+'A'-1))
	}
	return string(res)
}
