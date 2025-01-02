// 2279. 装满石头的背包的最大数量

package algorithm_2200

import (
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {

	for i := 0; i < len(capacity); i++ {
		capacity[i] -= rocks[i]
	}
	sort.Ints(capacity)
	var num int
	for i := 0; i < len(capacity) && additionalRocks > 0; i++ {
		if additionalRocks >= capacity[i] {
			num++
			additionalRocks -= capacity[i]
			continue
		}

		break
	}

	return num
}
