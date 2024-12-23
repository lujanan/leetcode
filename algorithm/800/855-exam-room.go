//在考场里，有 n 个座位排成一行，编号为 0 到 n - 1。
//
// 当学生进入考场后，他必须坐在离最近的人最远的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上
//。)
//
// 设计一个模拟所述考场的类。
//
// 实现 ExamRoom 类：
//
//
// ExamRoom(int n) 用座位的数量 n 初始化考场对象。
// int seat() 返回下一个学生将会入座的座位编号。
// void leave(int p) 指定坐在座位 p 的学生将离开教室。保证座位 p 上会有一位学生。
//
//
//
//
// 示例 1：
//
//
//输入：
//["ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"]
//[[10], [], [], [], [], [4], []]
//输出：
//[null, 0, 9, 4, 2, null, 5]
//解释：
//ExamRoom examRoom = new ExamRoom(10);
//examRoom.seat(); // 返回 0，房间里没有人，学生坐在 0 号座位。
//examRoom.seat(); // 返回 9，学生最后坐在 9 号座位。
//examRoom.seat(); // 返回 4，学生最后坐在 4 号座位。
//examRoom.seat(); // 返回 2，学生最后坐在 2 号座位。
//examRoom.leave(4);
//examRoom.seat(); // 返回 5，学生最后坐在 5 号座位。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
// 保证有学生正坐在座位 p 上。
// seat 和 leave 最多被调用 10⁴ 次。
//
//
// Related Topics 设计 有序集合 堆（优先队列） 👍 285 👎 0

package algorithm_800

// leetcode submit region begin(Prohibit modification and deletion)
type ExamRoom struct {
	MaxIdx   int
	Nums     []int
	SeatNums []int
}

func Constructor(n int) ExamRoom {

}

func (this *ExamRoom) Seat() int {

}

func (this *ExamRoom) Leave(p int) {

}



/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
//leetcode submit region end(Prohibit modification and deletion)
