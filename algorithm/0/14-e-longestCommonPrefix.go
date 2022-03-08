//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
// Related Topics 字符串 👍 2083 👎 0

package algorithm_0

func longestCommonPrefix(strs []string) string {
	var res []byte
	if len(strs) > 0 {
		str := strs[0]
		for i := range str {
			for j := 1; j < len(strs); j++ {
				if i >= len(strs[j]) || strs[j][i] != str[i] {
					return string(res)
				}
			}
			res = append(res, str[i])
		}
	}

	return string(res)
}
