// 3479.水果成篮III

package algorith_3400

func numOfUnplacedFruitsIII(fruits []int, baskets []int) int {
	// 构建线段树，节点为每个区间的最大值
	var ll = len(baskets)
	for i := 2; ; i <<= 1 { // 寻找构造完全二叉树的最小长度，可能会浪费空间
		if ll <= i {
			ll = i
			break
		}
	}

	var tree = make([]int, ll<<1-1)
	for i := 0; i < len(baskets); i++ {
		tree[i+ll-1] = baskets[i]
	}
	for i := ll - 1; i > 0; i-- {
		tree[i-1] = max(tree[i<<1-1], tree[i<<1])
	}

	var checkAndUpdateMaxTree = func(num int) bool {
		if tree[0] < num {
			return false
		}
		var hit = 1
		for i := 1; i < ll; {
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
