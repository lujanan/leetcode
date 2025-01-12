// 191. 位1的个数

package algorithm_100

// n & (n-1) 把n中最低位的1消掉
func hammingWeight(n int) int {
	var res int
	for ; n > 0; n &= n - 1 {
		res++
	}
	return res
}

func hammingWeightV2(n int) int {
	var res int
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
