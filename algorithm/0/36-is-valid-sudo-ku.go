//请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
//
// 注意：
//
//
// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
//
//
//
//
// 示例 1：
//
//
//输入：board =
//[["5","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//输出：true
//
//
// 示例 2：
//
//
//输入：board =
//[["8","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//输出：false
//解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无
//效的。
//
//
//
// 提示：
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字或者 '.'
//
// Related Topics 数组 哈希表 矩阵
// 👍 616 👎 0

package algorithm_0

// 位运算
func isValidSudoku(board [][]byte) bool {
	var row, col, rc = [9]int{}, [9]int{}, [9]int{}
	var n int
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] == '.' {
				continue
			}

			n = 1 << (board[y][x] - '1')
			if row[y]&n != 0 || col[x]&n != 0 || rc[y/3*3+x/3]&n != 0 {
				return false
			}
			row[y], col[x], rc[y/3*3+x/3] = row[y]|n, col[x]|n, rc[y/3*3+x/3]|n
		}
	}
	return true
}

// 暴力
func isValidSudoku1(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var checkRow = make(map[byte]bool)
		var checkCol = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			// 检查行
			if board[i][j] != '.' {
				if _, ok := checkRow[board[i][j]]; ok {
					return false
				}
				checkRow[board[i][j]] = true
			}
			// 检查列
			if board[j][i] != '.' {
				if _, ok := checkCol[board[j][i]]; ok {
					return false
				}
				checkCol[board[j][i]] = true
			}
		}
	}

	for n := 0; n < 9; n += 3 {
		for m := 0; m < 9; m += 3 {
			var checkRow = make(map[byte]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[i+n][j+m] != '.' {
						if _, ok := checkRow[board[i+n][j+m]]; ok {
							return false
						}
						checkRow[board[i+n][j+m]] = true
					}
				}
			}
		}
	}
	return true
}
