// 3297.统计重新排列后包含另一个字符串的子字符串数目I
//给你两个字符串 word1 和 word2 。
//
// 如果一个字符串 x 重新排列后，word2 是重排字符串的 前缀 ，那么我们称字符串 x 是 合法的 。
//
// 请你返回 word1 中 合法 子字符串 的数目。
//
//
//
// 示例 1：
//
//
// 输入：word1 = "bcca", word2 = "abc"
//
//
// 输出：1
//
// 解释：
//
// 唯一合法的子字符串是 "bcca" ，可以重新排列得到 "abcc" ，"abc" 是它的前缀。
//
// 示例 2：
//
//
// 输入：word1 = "abcabc", word2 = "abc"
//
//
// 输出：10
//
// 解释：
//
// 除了长度为 1 和 2 的所有子字符串都是合法的。
//
// 示例 3：
//
//
// 输入：word1 = "abcabc", word2 = "aaabc"
//
//
// 输出：0
//
//
//
// 解释：
//
//
// 1 <= word1.length <= 10⁵
// 1 <= word2.length <= 10⁴
// word1 和 word2 都只包含小写英文字母。
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 14 👎 0

package algorithm_3200

// leetcode submit region begin(Prohibit modification and deletion)

func validSubstringCount(word1 string, word2 string) int64 {
    diff := make([]int, 26)
    for _, c := range word2 {
        diff[c - 'a']--
    }
    cnt := 0
    for _, c := range diff {
        if c < 0 {
            cnt++
        }
    }
    var res int64
    l, r := 0, 0
    for l < len(word1) {
        for r < len(word1) && cnt > 0 {
            update(diff, int(word1[r] - 'a'), 1, &cnt)
            r++
        }
        if cnt == 0 {
            res += int64(len(word1) - r + 1)
        }
        update(diff, int(word1[l]-'a'), -1, &cnt)
        l++
    }

    return res
}

func update(diff []int, c, add int, cnt *int) {
    diff[c] += add
    if add == 1 && diff[c] == 0 {
        // 表明 diff[c] 由 -1 变为 0
        *cnt--
    } else if add == -1 && diff[c] == -1 {
        // 表明 diff[c] 由 0 变为 -1
        *cnt++
    }
}


// 简单滑动窗口，超时
func validSubstringCountV2(word1 string, word2 string) int64 {
	var (
		wordNum     = make([]int, 26)
		wordZeroMap = make(map[int]int)

		updateMap = func(idx ...int) {
			for _, i := range idx {
				if wordNum[i] >= 0 {
					delete(wordZeroMap, i)
				} else {
					wordZeroMap[i] = wordNum[i]
				}
			}
		}

		copyMap = func() map[int]int {
			var cpMap = make(map[int]int)
			for k, v := range wordZeroMap {
				cpMap[k] = v
			}
			return cpMap
		}
	)

	for _, v := range word2 {
		wordNum[v-'a']--
	}
	for i := 0; i < len(word1) && i < len(word2); i++ {
		wordNum[word1[i]-'a']++
	}
	for i := 0; i < len(wordNum); i++ {
		if wordNum[i] < 0 {
			wordZeroMap[i] = wordNum[i]
		}
	}

	var res int64
	for i, j := 0, len(word2)-1; i < len(word1) && j < len(word1); i, j = i+1, j+1 {
		if len(wordZeroMap) <= 0 {
			res += int64(len(word1) - j)
		} else {
			cpMap := copyMap()
			for k := j + 1; k < len(word1); k++ {
				var idx = int(word1[k] - 'a')
				if num, ok := cpMap[idx]; ok {
					num++
					if num >= 0 {
						delete(cpMap, idx)
					} else {
						cpMap[idx] = num
					}
				}
				if len(cpMap) <= 0 {
					res += int64(len(word1) - k)
					break
				}
			}
		}

		if i+1 < len(word1) && j+1 < len(word1) {
			wordNum[word1[i]-'a']--
			wordNum[word1[j+1]-'a']++
			updateMap(int(word1[i]-'a'), int(word1[j+1]-'a'))
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
