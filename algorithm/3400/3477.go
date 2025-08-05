// 3477.水果成篮II

package algorith_3400

func numOfUnplacedFruits(fruits []int, baskets []int) int {
    var ans int
	for _, f := range fruits {
		var found = false
		for i := 0; i < len(baskets); i++ {
			if f <= baskets[i] {
				found = true
				baskets[i] = 0
				break
			}
		}

		if !found {
			ans++
		}
	}
	return ans
}
