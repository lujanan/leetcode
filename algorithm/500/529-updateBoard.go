package algorithm_500

// 广度优先搜索
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	var dx = []int{-1, 1, 0, 0, -1, -1, 1, 1}
	var dy = []int{0, 0, -1, 1, -1, 1, 1, -1}

	var curr = [][]int{click}
	var next [][]int

loop:
	for _, point := range curr {
		if board[point[0]][point[1]] != 'E' {
			continue
		}

		var box int
		for k := 0; k < len(dx); k++ {
			var x, y = point[1] + dx[k], point[0] + dy[k]
			if x >= 0 && x < len(board[0]) && y >= 0 && y < len(board) && board[y][x] == 'M' {
				box++
			}
		}
		if box > 0 && box < 9 {
			board[point[0]][point[1]] = byte(box) + '0'
		} else {
			board[point[0]][point[1]] = 'B'
			for k := 0; k < len(dx); k++ {
				var x, y = point[1] + dx[k], point[0] + dy[k]
				if x >= 0 && x < len(board[0]) && y >= 0 && y < len(board) {
					next = append(next, []int{y, x})
				}
			}
		}
	}

	if len(next) > 0 {
		curr = curr[:0]
		curr = append(curr, next...)
		next = next[:0]
		goto loop
	}

	return board
}

// 深度优先搜索
func updateBoard1(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	var dx = []int{-1, 1, 0, 0, -1, -1, 1, 1}
	var dy = []int{0, 0, -1, 1, -1, 1, 1, -1}
	var fn func(click []int)
	fn = func(click []int) {
		if board[click[0]][click[1]] != 'E' {
			return
		}

		var box int
		for k := 0; k < len(dx); k++ {
			var x, y = click[1] + dx[k], click[0] + dy[k]
			if x >= 0 && x < len(board[0]) && y >= 0 && y < len(board) && board[y][x] == 'M' {
				box++
			}
		}

		if box > 0 && box < 9 {
			board[click[0]][click[1]] = byte(box) + '0'
		} else {
			board[click[0]][click[1]] = 'B'
			for k := 0; k < len(dx); k++ {
				var x, y = click[1] + dx[k], click[0] + dy[k]
				if x >= 0 && x < len(board[0]) && y >= 0 && y < len(board) {
					fn([]int{y, x})
				}
			}
		}
	}
	fn(click)

	return board
}
