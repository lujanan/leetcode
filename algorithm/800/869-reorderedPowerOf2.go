//给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。
//
// 如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。
//
//
//
//
//
//
// 示例 1：
//
// 输入：1
//输出：true
//
//
// 示例 2：
//
// 输入：10
//输出：false
//
//
// 示例 3：
//
// 输入：16
//输出：true
//
//
// 示例 4：
//
// 输入：24
//输出：false
//
//
// 示例 5：
//
// 输入：46
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= N <= 10^9
//
// Related Topics 数学 计数 枚举 排序 👍 120 👎 0

package algorithm_800

import (
	"fmt"
)

func reorderedPowerOf2(n int) bool {
	var arr []int
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	return false
}

func permutaionImpl(prefix []byte, s []byte, cur int) {
    if cur == len(s) {
        fmt.Println(string(prefix))
        return
    }

    for _, b := range s {
        exist := false
        for i := 0; i < cur; i++ {
            if prefix[i] == b {
                exist = true
                break
            }
        }

        if !exist {
            prefix[cur] = b
            permutaionImpl(prefix, s, cur+1)
        }

    }
}
