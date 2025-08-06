// 3479.水果成篮III

package algorith_3400

func numOfUnplacedFruitsIII(fruits []int, baskets []int) int {
	// 构建线段树，节点为每个区间的最大值
	var ll = len(baskets)
	var tree = make([]int, ll<<1+1)
	for i := ll; i < ll<<1; i++ {
		tree[i] = baskets[i-ll]
	}
	for i := ll; i > 0; i-- {
		tree[i-1] = max(tree[i<<1-1], tree[i<<1])
	}

	var checkAndUpdateMaxTree = func(num int) bool {
		if tree[0] < num {
			return false
		}
		var hit = 1
		for i := 1; i <= ll; {
			// 优先查左节点
			hit = i << 1
			i <<= 1
			if tree[i-1] < num {
				hit += 1
				i += 1
			}
		}

		// 置0，更新树
		tree[hit-1] = 0
		for i := hit >> 1; i > 0; i >>= 1 {
			tree[i-1] = max(tree[i<<1-1], tree[i<<1])
		}

		return true
	}

	var ans int
	for _, v := range fruits {
		if !checkAndUpdateMaxTree(v) {
			ans++
		}
	}
	return ans
}
