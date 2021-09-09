//给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
// 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
// 说明:
//
//
// 单词是指由非空格字符组成的字符序列。
// 每个单词的长度大于 0，小于等于 maxWidth。
// 输入单词数组 words 至少包含一个单词。
//
//
// 示例:
//
// 输入:
//words = ["This", "is", "an", "example", "of", "text", "justification."]
//maxWidth = 16
//输出:
//[
//   "This    is    an",
//   "example  of text",
//   "justification.  "
//]
//
//
// 示例 2:
//
// 输入:
//words = ["What","must","be","acknowledgment","shall","be"]
//maxWidth = 16
//输出:
//[
//  "What   must   be",
//  "acknowledgment  ",
//  "shall be        "
//]
//解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//     因为最后一行应为左对齐，而不是左右两端对齐。
//     第二行同样为左对齐，这是因为这行只包含一个单词。
//
//
// 示例 3:
//
// 输入:
//words = ["Science","is","what","we","understand","well","enough","to","explain
//",
//         "to","a","computer.","Art","is","everything","else","we","do"]
//maxWidth = 20
//输出:
//[
//  "Science  is  what we",
//  "understand      well",
//  "enough to explain to",
//  "a  computer.  Art is",
//  "everything  else  we",
//  "do                  "
//]
//
// Related Topics 字符串 模拟
// 👍 192 👎 0

package algorithm

import "strings"

func fullJustify(words []string, maxWidth int) (res []string) {
	var l, tmlLen = 0, 0
	for r := 0; r < len(words); r++ {
		tmlLen += len(words[r])
		if tmlLen == maxWidth || r+1 == len(words) || tmlLen+1+len(words[r+1]) > maxWidth {
			num := r - l + 1                          // 字符串数
			emptyNum := maxWidth - (tmlLen - num + 1) // 需加空格数
			var emptyArr []int
			if r+1 < len(words) {
				emptyArr = []int{emptyNum} // 两个单词以下时，仅需一个空格
				if num > 2 {
					emptyArr = make([]int, num-1)
					emptyNum1 := emptyNum / (num - 1)
					emptyNum2 := emptyNum % (num - 1)
					for i := 0; i < num-1; i++ {
						emptyArr[i] = emptyNum1
						if emptyNum2 > 0 {
							emptyArr[i]++
							emptyNum2--
						}
					}
				}
			}

			var str string
			for i := 0; i < num; i++ {
				if r+1 >= len(words) {
					// 最后一个，左对齐
					str += words[l+i]
					if i < num-1 {
						str += " "
					} else if len(str) < maxWidth {
						str += strings.Repeat(" ", maxWidth-len(str))
					}
				} else {
					str += words[l+i]
					if i < len(emptyArr) && emptyArr[i] > 0 {
						str += strings.Repeat(" ", emptyArr[i])
					}
				}
			}
			res = append(res, str)

			tmlLen = 0
			l = r + 1
			continue
		}
		tmlLen++
	}
	return
}
