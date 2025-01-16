// 3095. 或值至少 K 的最短子数组 I

package algorithm_3000

func minimumSubarrayLength(nums []int, k int) int {
	var arr = make([]int, 6)
	var cnt, min = 0, -1
	var update = func(num int, isAdd bool) {
		for i := 0; i < len(arr); i++ {
			if num&(1<<i) != 0 {
				if isAdd {
					arr[i]++
				} else {
					arr[i]--
					if arr[i] == 0 {
						cnt ^= (1 << i)
					}
				}
			}
		}
	}

	var i, j int
	for ; i < len(nums); i++ {
		for j < len(nums) && cnt < k {
			cnt |= nums[j]
			update(nums[j], true)
			if cnt >= k || j == len(nums)-1 {
				break
			}
			j++
		}

		if cnt >= k {
			if min == -1 || min > j-i+1 {
				min = j - i + 1
			}
		}
		update(nums[i], false)
		if cnt < k || j == i {
			j++
		}
	}
	return min
}
