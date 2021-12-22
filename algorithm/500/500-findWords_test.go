//给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。
//
// 美式键盘 中：
//
//
// 第一行由字符 "qwertyuiop" 组成。
// 第二行由字符 "asdfghjkl" 组成。
// 第三行由字符 "zxcvbnm" 组成。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：words = ["Hello","Alaska","Dad","Peace"]
//输出：["Alaska","Dad"]
//
//
// 示例 2：
//
//
//输入：words = ["omk"]
//输出：[]
//
//
// 示例 3：
//
//
//输入：words = ["adsdf","sfd"]
//输出：["adsdf","sfd"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 20
// 1 <= words[i].length <= 100
// words[i] 由英文字母（小写和大写字母）组成
//
// Related Topics 数组 哈希表 字符串 👍 193 👎 0

package algorithm_500

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{[]string{"Hello", "Alaska", "Dad", "Peace"}}, []string{"Alaska", "Dad"}},
		{"t2", args{[]string{"omk"}}, nil},
		{"t3", args{[]string{"adsdf", "sfd"}}, []string{"adsdf", "sfd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
