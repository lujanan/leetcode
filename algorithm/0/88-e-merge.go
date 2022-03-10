package algorithm_0

// 88. 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array/

func merge88(nums1 []int, m int, nums2 []int, n int) []int {
	for i, j, k := m-1, n-1, m+n-1; (i >= 0 || j >= 0) && k >= 0; k-- {
		if i >= 0 && j >= 0 {
			if nums1[i] >= nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else if i >= 0 {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	return nums1
}
