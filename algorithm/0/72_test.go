// 72.编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 3559 👎 0

package algorithm_0

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"t1", args{"horse", "ros"}, 3},
		// {"t2", args{"intention", "execution"}, 5},
		// {"t3", args{"horsqwuehqwie", "rosshdnfls"}, 11},
		{"t4", args{"zoologicoarchaeologist", "zoogeologist"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
