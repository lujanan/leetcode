package algorithm_600

import (
	"strconv"
)

func shoppingOffers(price []int, special [][]int, needs []int) (res int) {
	// 过滤掉购买价格高于单买的，或超数量无法购买的礼包
	var fSpecial [][]int
	for _, v := range special {
		var origin int
		for i := 0; i < len(needs); i++ {
			if v[i] > needs[i] {
				origin = -1
				break
			}
			origin += price[i] * v[i]
		}
		if origin >= 0 && origin > v[len(needs)] {
			fSpecial = append(fSpecial, v)
		}
	}
	// 存储重复计算的结果
	dp := make(map[string]int)
	return buySpecial(price, needs, fSpecial, dp)
}

func buySpecial(price []int, needs []int, special [][]int, dp map[string]int) int {
	var key, num = "", 0
	for _, v := range needs {
		key += strconv.Itoa(v)
		num += v
	}
	min, ok := dp[key]
	if ok {
		return min
	}
	if num <= 0 {
		return 0
	}

	min = 0
	for k, v := range needs {
		min += v * price[k]
	}

	for _, s := range special {
		var hit, nextNeeds = true, make([]int, len(needs))
		for i := 0; i < len(needs); i++ {
			if needs[i] < s[i] {
				hit = false
				break
			}
			nextNeeds[i] = needs[i] - s[i]
		}
		if hit {
			nextMin := buySpecial(price, nextNeeds, special, dp)
			if min > s[len(needs)]+nextMin {
				min = s[len(needs)] + nextMin
			}
		}
	}
	dp[key] = min
	return min
}
