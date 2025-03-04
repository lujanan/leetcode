//给你一个字符串 s ，如果可以将它分割成三个 非空 回文子字符串，那么返回 true ，否则返回 false 。
//
// 当一个字符串正着读和反着读是一模一样的，就称其为 回文字符串 。
//
//
//
// 示例 1：
//
//
//输入：s = "abcbdd"
//输出：true
//解释："abcbdd" = "a" + "bcb" + "dd"，三个子字符串都是回文的。
//
//
// 示例 2：
//
//
//输入：s = "bcbddxy"
//输出：false
//解释：s 没办法被分割成 3 个回文子字符串。
//
//
//
//
// 提示：
//
//
// 3 <= s.length <= 2000
// s 只包含小写英文字母。
//
//
// Related Topics 字符串 动态规划 👍 71 👎 0

package algorithm_1700

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func checkPartitioning(s string) bool {

    // dp[i][0] = dp[i+1][0] 

    var checkMap = make(map[string]bool)
    var checkFn func(s string, idx, num int) bool
    checkFn = func(s string, idx, num int) bool {
        var check bool


        checkMap[fmt.Sprintf("%d-%d", idx, num)] = check
        return check
    }
    return checkFn(s, 0, 2)
}
//leetcode submit region end(Prohibit modification and deletion)
