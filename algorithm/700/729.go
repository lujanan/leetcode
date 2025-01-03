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

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
type MyCalendar struct {
	BookList [][2]int
}

func ConstructorV2() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for i := 0; i < len(this.BookList); i++ {
		if endTime <= this.BookList[i][0] {
			break
		}
		if startTime >= this.BookList[i][1] {
			continue
		}
		return false
	}

	this.BookList = append(this.BookList, [2]int{startTime, endTime})
	sort.Slice(this.BookList, func(i int, j int) bool { return this.BookList[i][0] < this.BookList[j][0] })
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
//leetcode submit region end(Prohibit modification and deletion)
