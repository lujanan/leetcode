//给你一个数组 nums，我们可以将它按一个非负整数 k 进行轮调，这样可以使数组变为 [nums[k], nums[k + 1], ... nums[
//nums.length - 1], nums[0], nums[1], ..., nums[k-1]] 的形式。此后，任何值小于或等于其索引的项都可以记作一分。
//
//
// 例如，数组为 nums = [2,4,1,3,0]，我们按 k = 2 进行轮调后，它将变成 [1,3,0,2,4]。这将记为 3 分，因为 1 > 0
//[不计分]、3 > 1 [不计分]、0 <= 2 [计 1 分]、2 <= 3 [计 1 分]，4 <= 4 [计 1 分]。
//
//
// 在所有可能的轮调中，返回我们所能得到的最高分数对应的轮调下标 k 。如果有多个答案，返回满足条件的最小的下标 k 。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,1,4,0]
//输出：3
//解释：
//下面列出了每个 k 的得分：
//k = 0,  nums = [2,3,1,4,0],    score 2
//k = 1,  nums = [3,1,4,0,2],    score 3
//k = 2,  nums = [1,4,0,2,3],    score 3
//k = 3,  nums = [4,0,2,3,1],    score 4
//k = 4,  nums = [0,2,3,1,4],    score 3
//所以我们应当选择 k = 3，得分最高。
//
// 示例 2：
//
//
//输入：nums = [1,3,0,2,4]
//输出：0
//解释：
//nums 无论怎么变化总是有 3 分。
//所以我们将选择最小的 k，即 0。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] < nums.length
//
// Related Topics 数组 前缀和 👍 111 👎 0

package algorithm_700

func bestRotation(nums []int) int {
	var res, rmi, lmi int
	var n = len(nums)
	var kl = make([]int, len(nums)+1)
	for i := 0; i < n; i++ {
		rmi = (i - nums[i] + n) % n // 右边界
		// lmi = (i - nums[i] + n - (n - 1 - nums[i])) % n // 符合条件的有 n - 1 - nums[i] 个位置，左边界 = 右边界 - 数量
		lmi = (i + 1) % n

		kl[rmi+1]--
		kl[lmi]++
		if lmi >= rmi {
			kl[0]++
		}
	}

	for i := 1; i < n; i++ {
		kl[i] += kl[i-1]
		if kl[res] < kl[i] {
			res = i
		}
	}
	return res
}

// 差分数组，
// 题意：k++ 时，数组左移
// 计算 nums[i] == idx 时 k 的值，此时，下标 idx 继续右移，nums[i] < idx 恒成立，此时 k-- ,
// 当 k<0 时，k = n + k
// 引入数组 kl 记录 0~k-1 时的分数
// 差分数组，找到最小的 minK 和 最大的 maxK 的区间，kl[minK]+1, kl[maxK]-1 即可实现区间内的数值+1
// 而不影响其他区间的数值
// 注意：存在 k < 0 时存在两个区间；maxK 大于数组长度
func bestRotation2(nums []int) int {
	var res, rmi, lmi, nli int
	var n = len(nums)
	var kl = make([]int, len(nums))
	for i := 0; i < n; i++ {
		rmi = i - nums[i]
		if rmi < 0 {
			rmi += n
		}

		nli = n - nums[i]
		lmi = rmi - nli + 1
		if lmi >= 0 {
			kl[lmi]++
		} else {
			kl[0]++
			kl[n+lmi]++
		}

		if rmi+1 < n {
			kl[rmi+1]--
		}
	}

	for i := 1; i < n; i++ {
		kl[i] += kl[i-1]
		if kl[res] < kl[i] {
			res = i
		}
	}
	return res
}

// 超时 O(n^2)
func bestRotation1(nums []int) int {
	var res, mi, ti int
	var n = len(nums)
	for k := 0; k < n; k++ {
		ti = 0
		for i := 0; i < n; i++ {
			if nums[(k+i)%n] <= i {
				ti++
			}
		}
		if mi < ti {
			mi = ti
			res = k
		}
	}

	return res
}
