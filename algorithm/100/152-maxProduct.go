package algorithm_100

func maxProduct(nums []int) int {
	// 保留最大和最小，负负得正
	var mm, mn = nums[0], nums[0]
	var p0, p1 = 1, 1
	for i := 0; i < len(nums); i++ {
		p0, p1 = minn(p0*nums[i], p1*nums[i], nums[i]), maxn(p0*nums[i], p1*nums[i], nums[i])
		mm, mn = maxn(p1, mm), minn(p0, mn)
	}

	return mm
}

func maxn(list ...int) int {
	m := list[0]
	for i := 1; i < len(list); i++ {
		if m < list[i] {
			m = list[i]
		}
	}
	return m
}

func minn(list ...int) int {
	m := list[0]
	for i := 1; i < len(list); i++ {
		if m > list[i] {
			m = list[i]
		}
	}
	return m
}
