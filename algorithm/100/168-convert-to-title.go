//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
//
// 例如：
//
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//
//
// 示例 1：
//
//
//输入：columnNumber = 1
//输出："A"
//
//
// 示例 2：
//
//
//输入：columnNumber = 28
//输出："AB"
//
//
// 示例 3：
//
//
//输入：columnNumber = 701
//输出："ZY"
//
//
// 示例 4：
//
//
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
//
//
//
//
// 提示：
//
//
// 1 <= columnNumber <= 231 - 1
//
// Related Topics 数学 字符串
// 👍 457 👎 0

package algorithm_100

import "fmt"

func convertToTitle(columnNumber int) (res string) {
	for columnNumber > 0 {
		n := (columnNumber-1)%26 + 1
		res = fmt.Sprintf("%s%s", string('A'+byte(n-1)), res)
		columnNumber = (columnNumber - n) / 26
	}
	return
}
