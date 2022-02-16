//在大小为 n x n 的网格 grid 上，每个单元格都有一盏灯，最初灯都处于 关闭 状态。
//
// 给你一个由灯的位置组成的二维数组 lamps ，其中 lamps[i] = [rowi, coli] 表示 打开 位于 grid[rowi][coli]
//的灯。即便同一盏灯可能在 lamps 中多次列出，不会影响这盏灯处于 打开 状态。
//
// 当一盏灯处于打开状态，它将会照亮 自身所在单元格 以及同一 行 、同一 列 和两条 对角线 上的 所有其他单元格 。
//
// 另给你一个二维数组 queries ，其中 queries[j] = [rowj, colj] 。对于第 j 个查询，如果单元格 [rowj, colj]
// 是被照亮的，则查询结果为 1 ，否则为 0 。在第 j 次查询之后 [按照查询的顺序] ，关闭 位于单元格 grid[rowj][colj] 上及相邻 8 个
//方向上（与单元格 grid[rowi][coli] 共享角或边）的任何灯。
//
// 返回一个整数数组 ans 作为答案， ans[j] 应等于第 j 次查询 queries[j] 的结果，1 表示照亮，0 表示未照亮。
//
//
//
// 示例 1：
//
//
//输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
//输出：[1,0]
//解释：最初所有灯都是关闭的。在执行查询之前，打开位于 [0, 0] 和 [4, 4] 的灯。第 0 次查询检查 grid[1][1] 是否被照亮（蓝色方框）
//。该单元格被照亮，所以 ans[0] = 1 。然后，关闭红色方框中的所有灯。
//
//第 1 次查询检查 grid[1][0] 是否被照亮（蓝色方框）。该单元格没有被照亮，所以 ans[1] = 0 。然后，关闭红色矩形中的所有灯。
//
//
//
// 示例 2：
//
//
//输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
//输出：[1,1]
//
//
// 示例 3：
//
//
//输入：n = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
//输出：[1,1,0]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
// 0 <= lamps.length <= 20000
// 0 <= queries.length <= 20000
// lamps[i].length == 2
// 0 <= rowi, coli < n
// queries[j].length == 2
// 0 <= rowj, colj < n
//
// Related Topics 数组 哈希表 👍 132 👎 0

package algorithm_1000

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	var res []int
	// 顺时针转角度 0 45 90 135 180 225 270 315
	var dx = []int{0, 1, 1, 1, 0, -1, -1, -1}
	var dy = []int{-1, -1, 0, 1, 1, 1, 0, -1}

	// 灯网格
	var martx = make(map[[2]int]int)
	for _, l := range lamps {
		martx[[2]int{l[0], l[1]}] = 1
	}

	// 搜索是否亮
	var fn = func(y, x int) bool {
		if _, ok := martx[[2]int{y, x}]; ok {
			return true
		}
		// 是否被照亮
		for i := 0; i < 8; i++ {
			ny, nx := y, x
			for {
				ty, tx := ny+dy[i], nx+dx[i]
				if ty < 0 || ty >= n || tx < 0 || tx >= n {
					break
				}
				if _, ok := martx[[2]int{ty, tx}]; ok {
					return true
				}
				ny, nx = ty, tx
			}
		}
		return false
	}

	for _, q := range queries {
		if !fn(q[0], q[1]) {
			res = append(res, 0)
			continue
		}

		res = append(res, 1)
		delete(martx, [2]int{q[0], q[1]})
		for i := 0; i < 8; i++ {
			delete(martx, [2]int{q[0] + dy[i], q[1] + dx[i]})
		}
	}
	return res
}

// 依然内存溢出
func gridIllumination2(n int, lamps [][]int, queries [][]int) []int {
	var res []int
	// 顺时针转角度 0 45 90 135 180 225 270 315
	var dx = []int{0, 1, 1, 1, 0, -1, -1, -1}
	var dy = []int{-1, -1, 0, 1, 1, 1, 0, -1}

	// 灯网格
	var martx = make(map[[2]int]int)

	var fn = func(x, y, d, add int) {
		for {
			nx, ny := x+dx[d], y+dy[d]
			if nx < 0 || nx >= n || ny < 0 || ny >= n {
				break
			}
			martx[[2]int{ny, nx}] += add
			if martx[[2]int{ny, nx}] < 0 {
				delete(martx, [2]int{ny, nx})
			}
			x, y = nx, ny
		}
	}

	for _, l := range lamps {
		if martx[[2]int{l[0], l[1]}]%2 == 0 {
			// 点灯 数值+1 (不可重复点灯)
			martx[[2]int{l[0], l[1]}] += 1
			// 点亮各个方向 数值+2 (可被多个灯重复点亮)
			for i := 0; i < 8; i++ {
				fn(l[1], l[0], i, 2)
			}
		}
	}

	for _, q := range queries {
		// 不亮
		if martx[[2]int{q[0], q[1]}] <= 0 {
			res = append(res, 0)
			continue
		}
		// 亮着
		res = append(res, 1)
		// 关灯
		if martx[[2]int{q[0], q[1]}]%2 == 1 {
			martx[[2]int{q[0], q[1]}] -= 1
			for i := 0; i < 8; i++ {
				// 还原照亮的地方
				fn(q[1], q[0], i, -2)
			}
		}
		// 检查周边8个格子
		for i := 0; i < 8; i++ {
			y, x := q[0]+dy[i], q[1]+dx[i]
			if x < 0 || x >= n || y < 0 || y >= n {
				continue
			}
			if martx[[2]int{y, x}]%2 == 1 {
				martx[[2]int{y, x}] -= 1
				for k := 0; k < 8; k++ {
					// 还原照亮的地方
					fn(x, y, k, -2)
				}
			}
		}
	}
	return res
}

// 内存溢出
func gridIllumination1(n int, lamps [][]int, queries [][]int) []int {
	var res []int
	// 顺时针转角度 0 45 90 135 180 225 270 315
	var dx = []int{0, 1, 1, 1, 0, -1, -1, -1}
	var dy = []int{-1, -1, 0, 1, 1, 1, 0, -1}

	// 灯网格
	var martx = make([][]int, n)
	for i := 0; i < n; i++ {
		martx[i] = make([]int, n)
	}

	var fn func(x, y, d, add int)
	fn = func(x, y, d, add int) {
		nx, ny := x+dx[d], y+dy[d]
		if nx < 0 || nx >= n || ny < 0 || ny >= n {
			return
		}
		martx[ny][nx] += add
		if martx[ny][nx] < 0 {
			martx[ny][nx] = 0
		}
		fn(nx, ny, d, add)
	}

	for _, l := range lamps {
		if martx[l[0]][l[1]]%2 == 0 {
			// 点灯 数值+1 (不可重复点灯)
			martx[l[0]][l[1]] += 1
			// 点亮各个方向 数值+2 (可被多个灯重复点亮)
			for i := 0; i < 8; i++ {
				fn(l[1], l[0], i, 2)
			}
		}
	}

	for _, q := range queries {
		// 不亮
		if martx[q[0]][q[1]] <= 0 {
			res = append(res, 0)
			continue
		}
		// 亮着
		res = append(res, 1)
		// 关灯
		if martx[q[0]][q[1]]%2 == 1 {
			martx[q[0]][q[1]] -= 1
			for i := 0; i < 8; i++ {
				// 还原照亮的地方
				fn(q[1], q[0], i, -2)
			}
		}
		// 检查周边8个格子
		for i := 0; i < 8; i++ {
			y, x := q[0]+dy[i], q[1]+dx[i]
			if x < 0 || x >= n || y < 0 || y >= n {
				continue
			}
			if martx[y][x]%2 == 1 {
				martx[y][x] -= 1
				for k := 0; k < 8; k++ {
					// 还原照亮的地方
					fn(x, y, k, -2)
				}
			}
		}
	}

	return res
}
