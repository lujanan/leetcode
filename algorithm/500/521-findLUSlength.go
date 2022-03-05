package algorithm_500

// 521. 最长特殊序列 Ⅰ
// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/

func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }
    al, bl := len(a), len(b)
    if al > bl {
        return al
    }
    return bl
}