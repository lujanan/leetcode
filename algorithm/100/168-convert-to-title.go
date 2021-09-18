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

func convertToTitle(columnNumber int) string {
	var charMap = map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "N",
		14: "M",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}

	var numArr = []int{1}
	for i := 1; ; i++ {
		num := 26 * numArr[i-1]
		numArr = append(numArr, num)
		if num > columnNumber {
			break
		}
	}
	var res string
	for i := len(numArr) - 1; i >= 0; i-- {
		n := columnNumber / numArr[i]
		res = fmt.Sprintf("%s%s", res, charMap[n])
		columnNumber = columnNumber - n*numArr[i]
	}
	return res
}
