// 2080.区间内查询数字的频率

package algorithm_2000

type RangeFreqQuery struct {
	ArrMap map[int][]int
}

func ConstructorV2(arr []int) RangeFreqQuery {
	var aMap = make(map[int][]int)
	for i, num := range arr {
		aMap[num] = append(aMap[num], i)
	}
	return RangeFreqQuery{ArrMap: aMap}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	arr, ok := this.ArrMap[value]
	if !ok || len(arr) < 1 {
		return 0
	}

	var binarySearch = func(num, l, r int, isMax bool) int {
		for l <= r {
			var m = (l + r) >> 1
			if arr[m] == num {
				return m
			}

			if arr[m] < num {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if isMax {
			return l
		}
		return r
	}

	var l = binarySearch(left, 0, len(arr)-1, true)
	var r = binarySearch(right, 0, len(arr)-1, false)
	return r - l + 1
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
