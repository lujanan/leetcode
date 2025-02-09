// 732. 我的日程安排表 III

package algorithm_700

// 1.动态线段树
// 2.懒标志，0~10e9区间内，如果1~10都为1，那直接标志为1，不需要标志10个都为1

type MyCalendarThree struct {
	num, lazy map[int]int
}

func ConstructorV4() MyCalendarThree {
	return MyCalendarThree{
		num:  make(map[int]int),
		lazy: make(map[int]int),
	}
}

func (this MyCalendarThree) update(start, end, l, r, idx int) {
	if start > r || end < l {
		return
	}

	if start <= l && r <= end {
		this.num[idx]++
		this.lazy[idx]++
		return
	}

	mid := (l + r) >> 1
	this.update(start, end, l, mid, idx<<1)
	this.update(start, end, mid+1, r, idx<<1+1)
	this.num[idx] = this.lazy[idx] + max(this.num[idx<<1], this.num[idx<<1+1])
}

func (this MyCalendarThree) Book(start, end int) int {
	this.update(start, end-1, 0, 10e9, 1)
	return this.num[1]
}
