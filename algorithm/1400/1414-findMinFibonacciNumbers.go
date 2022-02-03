package algorithm_1400

func findMinFibonacciNumbers(k int) (res int) {
	var arr = []int{1, 1}
	for arr[len(arr)-1] < k {
		arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
	}
	for i := len(arr) - 1; i >= 0 && k > 0; i-- {
		if k >= arr[i] {
			k -= arr[i]
			res++
		}
	}
	return
}
