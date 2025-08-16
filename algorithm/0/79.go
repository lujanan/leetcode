// 79.单词搜索

package algorithm_0

func exist(board [][]byte, word string) bool {
	var direct = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var checkBoard = make(map[[2]int]bool)
	var ll, m, n = len(word), len(board), len(board[0])
	
	var dfs func(x, y, d, idx int) bool
	dfs = func(x, y, d, idx int) bool {
		if board[x][y] != word[idx] || checkBoard[[2]int{x, y}] {
			return false
		}

		if board[x][y] == word[idx] && idx == ll-1 {
			return true
		}

		checkBoard[[2]int{x, y}] = true
		for i := 0; i < 4; i++ {
			if x+direct[(d+i)%4][0] < 0 || x+direct[(d+i)%4][0] >= m ||
				y+direct[(d+i)%4][1] < 0 || y+direct[(d+i)%4][1] >= n ||
				idx+1 >= ll {
				continue
			}

			if dfs(x+direct[(d+i)%4][0], y+direct[(d+i)%4][1], (d+i)%4, idx+1) {
				return true
			}
		}
		checkBoard[[2]int{x, y}] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0, 0) {
				return true
			}
		}
	}

	return false
}
