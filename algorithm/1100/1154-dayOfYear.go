//给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。
//
// 通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。
//
//
//
// 示例 1：
//
//
//输入：date = "2019-01-09"
//输出：9
//
//
// 示例 2：
//
//
//输入：date = "2019-02-10"
//输出：41
//
//
// 示例 3：
//
//
//输入：date = "2003-03-01"
//输出：60
//
//
// 示例 4：
//
//
//输入：date = "2004-03-01"
//输出：61
//
//
//
// 提示：
//
//
// date.length == 10
// date[4] == date[7] == '-'，其他的 date[i] 都是数字
// date 表示的范围从 1900 年 1 月 1 日至 2019 年 12 月 31 日
//
// Related Topics 数学 字符串 👍 84 👎 0

package algorithm_1100

import "time"

func dayOfYear(date string) int {
	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	var days = day.Day()
	for i := 1; i < int(day.Month()); i++ {
		switch i {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			days += 31
		case 4:
			fallthrough
		case 6:
			fallthrough
		case 9:
			fallthrough
		case 11:
			days += 30
		case 2:
			days += 28
			if (day.Year()%4 == 0 && day.Year()%100 != 0) ||
				(day.Year()%100 == 0 && day.Year()%400 == 0) {
				days++
			}

		}

	}
	return days
}
