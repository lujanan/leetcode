//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
// 示例 1：
//
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//
//
// 示例 2：
//
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//
//
// 提示：
//
//
// 2 <= timePoints.length <= 2 * 10⁴
// timePoints[i] 格式为 "HH:MM"
//
// Related Topics 数组 数学 字符串 排序 👍 164 👎 0

package algorithm_500

import (
	"sort"
	"strconv"
	"strings"
)

// 暴力
func findMinDifference(timePoints []string) (res int) {
	var minute []int
	for i := 0; i < len(timePoints); i++ {
		tp := strings.Split(timePoints[i], ":")
		hour, _ := strconv.ParseInt(tp[0], 10, 64)
		min, _ := strconv.ParseInt(tp[1], 10, 64)
		minute = append(minute, int(hour*60+min))
	}
	sort.Ints(minute)
	minute = append(minute, minute[0]+24*60)
	res = minute[1] - minute[0]
	for i := 2; i < len(minute); i++ {
		if res > minute[i]-minute[i-1] {
			res = minute[i] - minute[i-1]
		}
	}
	return
}
