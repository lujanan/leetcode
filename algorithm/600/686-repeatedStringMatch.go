//给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
//
// 注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。
//
//
//
// 示例 1：
//
// 输入：a = "abcd", b = "cdabcdab"
//输出：3
//解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。
//
//
// 示例 2：
//
// 输入：a = "a", b = "aa"
//输出：2
//
//
// 示例 3：
//
// 输入：a = "a", b = "a"
//输出：1
//
//
// 示例 4：
//
// 输入：a = "abc", b = "wxyz"
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= a.length <= 10⁴
// 1 <= b.length <= 10⁴
// a 和 b 由小写英文字母组成
//
// Related Topics 字符串 字符串匹配 👍 194 👎 0

package algorithm_600

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

func repeatedStringMatch(a string, b string) int {
	an, bn := len(a), len(b)
	index := strStr(a, b)
	if index == -1 {
		return -1
	}
	if an-index >= bn {
		return 1
	}
	return (bn+index-an-1)/an + 2

}

func strStr(a string, b string) int {
	n, m := len(a), len(b)
	if m == 0 {
		return 0
	}
	var k1 int = 1000000000 + 7
	var k2 int = 1337
	rand.Seed(time.Now().Unix())
	var kMod1 int64 = int64(rand.Intn(k1)) + int64(k1)
	var kMod2 int64 = int64(rand.Intn(k2)) + int64(k2)
	var hash_b int64 = 0
	for i := 0; i < m; i++ {
		hash_b = (hash_b*kMod2 + int64(b[i])) % kMod1
	}
	var hash_a int64 = 0
	var extra int64 = 1
	for i := 0; i < m-1; i++ {
		hash_a = (hash_a*kMod2 + int64(a[i%n])) % kMod1
		extra = (extra * kMod2) % kMod1
	}
	for i := m - 1; (i - m + 1) < n; i++ {
		hash_a = (hash_a*kMod2 + int64(a[i%n])) % kMod1
		if hash_a == hash_b {
			return i - m + 1
		}
		hash_a = (hash_a - extra*int64(a[(i-m+1)%n])) % kMod1
		hash_a = (hash_a + kMod1) % kMod1
	}
	return -1
}

// 暴力解
func repeatedStringMatch0(a string, b string) int {
	na, nb := len(a), len(b)
	res := 1
	if na < nb {
		res = int(math.Ceil(float64(nb) / float64(na)))
	}
	if strings.Contains(strings.Repeat(a, res), b) {
		return res
	}
	if strings.Contains(strings.Repeat(a, res+1), b) {
		return res + 1
	}
	return -1
}