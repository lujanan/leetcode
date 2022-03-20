package algorithm_competition

func countHillValley(nums []int) int {
	var res int
	var n = len(nums)
	for i := 1; i < n-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > nums[i-1] {
			var r = i + 1
			for r < n && nums[i] == nums[r] {
				r++
			}
			if r < n && nums[r] < nums[i] {
				// 峰
				res++
				continue
			}
		}

		if nums[i] < nums[i-1] {
			var r = i + 1
			for r < n && nums[i] == nums[r] {
				r++
			}
			if r < n && nums[r] > nums[i] {
				// 谷
				res++
				continue
			}
		}
	}

	return res
}

func countCollisions(directions string) int {
	var res int
	var n = len(directions)
	var char = []byte(directions)

	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx < 0 || idx >= n {
			return 0
		}

		switch char[idx] {
		case 'R':
			if idx+1 >= n {
				return 0
			}

			if char[idx+1] == 'L' {
				res += 2
				char[idx], char[idx+1] = 'S', 'S'
				return 1
			} else if char[idx+1] == 'S' {
				res++
				char[idx], char[idx+1] = 'S', 'S'
				return 1
			} else {
				res += dfs(idx + 1)
			}

		case 'L':
			if idx-1 < 0 {
				return 0
			}
			if char[idx-1] == 'R' {
				res += 2
				char[idx], char[idx-1] = 'S', 'S'
				return 1
			} else if char[idx-1] == 'S' {
				res++
				char[idx], char[idx-1] = 'S', 'S'
				return 1
			} else {
				res += dfs(idx - 1)
			}
		}
		return 0
	}

	for i := 0; i < n; i++ {
		dfs(i)
	}
	return res
}
