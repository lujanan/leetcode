// 731.我的日程安排表II
//实现一个程序来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
//
// 当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生 三重预订。
//
// 事件能够用一对整数 startTime 和 endTime 表示，在一个半开区间的时间 [startTime, endTime) 上预定。实数 x 的范围
//为 startTime <= x < endTime。
//
// 实现 MyCalendarTwo 类：
//
//
// MyCalendarTwo() 初始化日历对象。
// boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否
//则，返回 false 并且不要将该日程安排添加到日历中。
//
//
//
//
// 示例 1：
//
//
//输入：
//["MyCalendarTwo", "book", "book", "book", "book", "book", "book"]
//[[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
//输出：
//[null, true, true, true, false, true, true]
//
//解释：
//MyCalendarTwo myCalendarTwo = new MyCalendarTwo();
//myCalendarTwo.book(10, 20); // 返回 True，能够预定该日程。
//myCalendarTwo.book(50, 60); // 返回 True，能够预定该日程。
//myCalendarTwo.book(10, 40); // 返回 True，该日程能够被重复预定。
//myCalendarTwo.book(5, 15);  // 返回 False，该日程导致了三重预定，所以不能预定。
//myCalendarTwo.book(5, 10); // 返回 True，能够预定该日程，因为它不使用已经双重预订的时间 10。
//myCalendarTwo.book(25, 55); // 返回 True，能够预定该日程，因为时间段 [25, 40) 将被第三个日程重复预定，时间段
//[40, 50) 将被单独预定，而时间段 [50, 55) 将被第二个日程重复预定。
//
//
//
//
// 提示：
//
//
// 0 <= start < end <= 10⁹
// 最多调用 book 1000 次。
//
//
// Related Topics 设计 线段树 数组 二分查找 有序集合 前缀和 👍 263 👎 0

package algorithm_700

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
type MyCalendarTwo struct {
	BookOne, BookTwo []int
	num, lasy        map[int]int
}

func ConstructorV3() MyCalendarTwo {
	return MyCalendarTwo{
		num:  make(map[int]int),
		lasy: make(map[int]int),
	}
}

func (this MyCalendarTwo) update(start, end, l, r, idx int) {
	if start > r || end < l {
		return
	}

	if start <= l && r <= end {
		this.num[idx]++
		this.lasy[idx]++
		return
	}

	mid := (l + r) >> 1
	this.update(start, end, l, mid, idx<<1)
	this.update(start, end, mid+1, r, idx<<1+1)
	this.num[idx] = this.lasy[idx] + int(math.Max(float64(this.num[idx<<1]), float64(this.num[idx<<1+1])))
}

func (this MyCalendarTwo) Book(start, end int) bool {

}

func (this MyCalendarTwo) BookV2(startTime int, endTime int) bool {
	for i := 1; i < len(this.BookTwo); i = i + 2 {
		if !(startTime >= this.BookTwo[i] || endTime <= this.BookTwo[i-1]) {
			return false
		}
	}

	for i := 1; i < len(this.BookOne); i = i + 2 {
		if startTime >= this.BookOne[i] || endTime <= this.BookOne[i-1] {
			continue
		}
		this.BookTwo = append(this.BookTwo,
			int(math.Max(float64(startTime), float64(this.BookOne[i-1]))),
			int(math.Min(float64(endTime), float64(this.BookOne[i]))))
	}

	this.BookOne = append(this.BookOne, startTime, endTime)
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
//leetcode submit region end(Prohibit modification and deletion)
