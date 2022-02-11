package algorithm_500

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	var nMap = map[int]byte{
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
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
			board[click[0]][click[1]] = nMap[box]
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
