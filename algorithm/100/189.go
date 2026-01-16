// 189. 轮转数组
//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁵
//
//
//
//
// 进阶：
//
//
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//
//
// Related Topics 数组 数学 双指针 👍 2284 👎 0

package algorithm_100

// leetcode submit region begin(Prohibit modification and deletion)

func rotateV4(nums []int, k int) []int {
	var ln = len(nums)
	k %= ln
	if k < 1 || k == ln || ln <= 1 {
		return nums
	}

	for i, j := 0, ln - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	for i, j := 0, k - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	for i, j := k, ln - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

// 翻转数组
// []int{1, 2, 3, 4, 5, 6, 7} k=3
// 1. 先整体翻转数组 []int{7, 6, 5, 4, 3, 2, 1}
// 2. 翻转前k个元素 []int{5, 6, 7, 4, 3, 2, 1}
// 3. 翻转后n-k个元素 []int{5, 6, 7, 1, 2, 3, 4}
// 时间复杂度O(n)，空间复杂度O(1)
func rotate(nums []int, k int) []int {
	var nl = len(nums)
	k %= nl
	if k == 0 || k == nl || nl <= 1 {
		return nums
	}

	for i, j := 0, nl-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]	
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, nl-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

// 时间复杂度O(n)，空间复杂度O(1)
// 通过最大公约数来计算循环次数
// 1. 从第一个元素开始，每次移动k个位置，直到回到起始位置
func rotateV1(nums []int, k int) []int {
	var nl = len(nums)
	k %= nl
	var gcb func(a, b int) int
	gcb = func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n int
	for st, cnt := 0, gcb(nl, k); st < cnt; st++ {
		var tmp = nums[st]
		for i := st + k; ; i += k {
			nums[i%nl], tmp = tmp, nums[i%nl]
			n++
			if i%nl == st || n > nl {
				break
			}
		}
	}
	return nums
}

// 时间复杂度O(n)，空间复杂度O(k)
// 通过额外数组来存储需要移动的元素
func rotateV2(nums []int, k int) []int {
	var nl = len(nums)
	k %= nl
	if k == 0 || k == nl || nl <= 1 {
		return nums
	}

	var tmpArr = make([]int, k)
	for i := 0; i < k; i++ {
		tmpArr[i] = nums[i]
	}
	for i := nl - 1; i >= k; i-- {
		nums[(i+k)%nl] = nums[i]
	}
	for i := 0; i < len(tmpArr); i++ {
		nums[(i+k)%nl] = tmpArr[i]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
