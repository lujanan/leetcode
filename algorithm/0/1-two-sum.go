package algorithm_0

func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int)
	for i, v := range nums {
		if j, ok := numMap[target-v]; ok {
			return []int{j, i}
		}
		numMap[v] = i
	}
	return nil
}
