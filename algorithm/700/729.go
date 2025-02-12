// 729. 我的日程安排表I
//实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
//
// 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
//
// 日程可以用一对整数 startTime 和 endTime 表示，这里的时间是半开区间，即 [startTime, endTime), 实数 x 的范围为
//， startTime <= x < endTime 。
//
// 实现 MyCalendar 类：
//
//
// MyCalendar() 初始化日历对象。
// boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。
//否则，返回 false 并且不要将该日程安排添加到日历中。
//
//
//
//
// 示例：
//
//
//输入：
//["MyCalendar", "book", "book", "book"]
//[[], [10, 20], [15, 25], [20, 30]]
//输出：
//[null, true, false, true]
//
//解释：
//MyCalendar myCalendar = new MyCalendar();
//myCalendar.book(10, 20); // return True
//myCalendar.book(15, 25); // return False ，这个日程安排不能添加到日历中，因为时间 15 已经被另一个日程安排预订了
//。
//myCalendar.book(20, 30); // return True ，这个日程安排可以添加到日历中，因为第一个日程安排预订的每个时间都小于 20
// ，且不包含时间 20 。
//
//
//
// 提示：
//
//
// 0 <= start < end <= 10⁹
// 每个测试用例，调用 book 方法的次数最多不超过 1000 次。
//
//
// Related Topics 设计 线段树 数组 二分查找 有序集合 👍 322 👎 0

package algorithm_700

// leetcode submit region begin(Prohibit modification and deletion)
type MyCalendar struct {
	BookList  []int
	num, lasy map[int]int
}

func ConstructorV2() MyCalendar {
	return MyCalendar{
		num:  make(map[int]int),
		lasy: make(map[int]int),
	}
}

func (this MyCalendar) update(start, end, l, r, idx int) {
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
	this.num[idx] = this.lasy[idx] + max(this.num[idx<<1], this.num[idx<<1+1])
}

func (this MyCalendar) query(start, end, l, r, idx int) int {
	if start > r || end < l {
		return 0
	}

	if start <= l && r <= end {
		return this.num[idx]
	}

	mid := (l + r) >> 1
	left := this.query(start, end, l, mid, idx<<1)
	right := this.query(start, end, mid+1, r, idx<<1+1)
	return this.lasy[idx] + max(left, right)
}

func (this MyCalendar) Book(startTime int, endTime int) bool {
	if this.query(startTime, endTime-1, 0, 10e9, 1) > 0 {
		return false
	}
	this.update(startTime, endTime-1, 0, 10e9, 1)
	return true
}

func (this MyCalendar) BookV1(startTime int, endTime int) bool {
	for i := 1; i < len(this.BookList); i = i + 2 {
		if !(startTime >= this.BookList[i] || endTime <= this.BookList[i-1]) {
			return false
		}
	}
	this.BookList = append(this.BookList, startTime, endTime)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
//leetcode submit region end(Prohibit modification and deletion)
