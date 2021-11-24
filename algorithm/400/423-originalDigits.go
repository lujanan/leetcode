//给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
//
//
//
// 示例 1：
//
//
//输入：s = "owoztneoer"
//输出："012"
//
//
// 示例 2：
//
//
//输入：s = "fviefuro"
//输出："45"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
// s 保证是一个符合题目要求的字符串
//
// Related Topics 哈希表 数学 字符串 👍 145 👎 0

package algorithm_400

import (
	"strconv"
	"strings"
)

func originalDigits(s string) (res string) {
	var nums = make([]int, 10)
	var sMap = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = 1
		} else {
			sMap[s[i]]++
		}
	}

	if _, ok := sMap['z']; ok {
		nums[0] = sMap['z']
	}
	if _, ok := sMap['w']; ok {
		nums[2] = sMap['w']
	}
	if _, ok := sMap['u']; ok {
		nums[4] = sMap['u']
	}
	if _, ok := sMap['u']; ok {
		nums[4] = sMap['u']
	}
	if _, ok := sMap['x']; ok {
		nums[6] = sMap['x']
	}
	if _, ok := sMap['g']; ok {
		nums[8] = sMap['g']
	}
	if _, ok := sMap['h']; ok {
		nums[3] = sMap['h'] - nums[8]
	}
	if _, ok := sMap['f']; ok {
		nums[5] = sMap['f'] - nums[4]
	}
	if _, ok := sMap['s']; ok {
		nums[7] = sMap['s'] - nums[6]
	}
	if _, ok := sMap['o']; ok {
		nums[1] = sMap['o'] - nums[0] - nums[2] - nums[4]
	}
	if _, ok := sMap['i']; ok {
		nums[9] = sMap['i'] - nums[5] - nums[6] - nums[8]
	}

	for i := 0; i < len(nums); i++ {
		res += strings.Repeat(strconv.Itoa(i), nums[i])
	}
	return
}
