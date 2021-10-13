//给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
//
//
// answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
// answer[i] == "Fizz" 如果 i 是 3 的倍数。
// answer[i] == "Buzz" 如果 i 是 5 的倍数。
// answer[i] == i 如果上述条件全不满足。
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["1","2","Fizz"]
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：["1","2","Fizz","4","Buzz"]
//
//
// 示例 3：
//
//
//输入：n = 15
//输出：["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","1
//4","FizzBuzz"]
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
//
// Related Topics 数学 字符串 模拟 👍 142 👎 0

package algorithm_400

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", args{3}, []string{"1", "2", "Fizz"}},
		{"t2", args{5}, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"t3", args{15}, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := fizzBuzz(tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("fizzBuzz() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
