package algorithm_0

// 77. 组合
// https://leetcode-cn.com/problems/combinations/

func combine(n int, k int) [][]int {
	var res [][]int
	var visited = make([]int, n+1)
	var arr []int

	var fn func(num int)
	fn = func(num int) {
		if num == k {
			res = append(res, append([]int{}, arr...))
			return
		}

		for i := 1; i <= n; i++ {
			if visited[i] == 1 || (len(arr) > 0 && arr[len(arr)-1] > i) {
				continue
			}

			visited[i] = 1
			arr = append(arr, i)
			fn(num + 1)
			visited[i] = 0
			arr = arr[:len(arr)-1]
		}

	}
	fn(0)
	return res
}
