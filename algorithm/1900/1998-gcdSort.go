package algorithm_1900

func gcdSort(nums []int) (res bool) {
	res = true
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			g := gcd(nums[j], nums[j-1])
			if g > 1 {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				return false
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
