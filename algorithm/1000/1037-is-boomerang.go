//回旋镖定义为一组三个点，这些点各不相同且不在一条直线上。
//
// 给出平面上三个点组成的列表，判断这些点是否可以构成回旋镖。
//
//
//
// 示例 1：
//
// 输入：[[1,1],[2,3],[3,2]]
//输出：true
//
//
// 示例 2：
//
// 输入：[[1,1],[2,2],[3,3]]
//输出：false
//
//
//
// 提示：
//
//
// points.length == 3
// points[i].length == 2
// 0 <= points[i][j] <= 100
//
// Related Topics 几何 数学
// 👍 28 👎 0

package algorithm_1000

func isBoomerang(points [][]int) bool {
	// 直线斜率k: (y1-y2)=k*(x1-x2)
	return (points[0][1]-points[1][1])*(points[1][0]-points[2][0]) != (points[0][0]-points[1][0])*(points[1][1]-points[2][1])
}
