// 79.单词搜索

package algorithm_0

func exist(board [][]byte, word string) bool {
	var direct = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var checkBoard = make(map[[2]int]bool)
	var ll = len(word)
	var check bool
	var dfs func(x, y, d, idx int) bool
	dfs = func(x, y, d, idx int) bool {
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || idx >= ll {
			return false
		}
		if checkBoard[[2]int{x, y}] {
			return false
		}

		if board[x][y] == word[idx] && idx == ll-1 {
			return true
		}

		if board[x][y] == word[idx] {
			checkBoard[[2]int{x, y}] = true
			for i := 0; i < 4; i++ {
				check = dfs(x+direct[(d+i)%4][0], y+direct[(d+i)%4][1], (d+i)%4, idx+1)
				if check {
					break
				}
			}
		}

		checkBoard[[2]int{x, y}] = false
		return check
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			if board[i][j] == word[0] && ll == 1 {
				return true
			}
			checkBoard[[2]int{i, j}] = true

			for d := 0; d < 4; d++ {
				if dfs(i+direct[d][0], j+direct[d][1], d, 1) {
					return true
				}
			}
			checkBoard[[2]int{i, j}] = false
		}
	}

	return false
}
