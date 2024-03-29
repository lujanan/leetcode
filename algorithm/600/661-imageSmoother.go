//包含整数的二维矩阵 M 表示一个图片的灰度。你需要设计一个平滑器来让每一个单元的灰度成为平均灰度 (向下舍入) ，平均灰度的计算是周围的8个单元和它本身的值
//求平均，如果周围的单元格不足八个，则尽可能多的利用它们。
//
// 示例 1:
//
//
//输入:
//[[1,1,1],
// [1,0,1],
// [1,1,1]]
//输出:
//[[0, 0, 0],
// [0, 0, 0],
// [0, 0, 0]]
//解释:
//对于点 (0,0), (0,2), (2,0), (2,2): 平均(3/4) = 平均(0.75) = 0
//对于点 (0,1), (1,0), (1,2), (2,1): 平均(5/6) = 平均(0.83333333) = 0
//对于点 (1,1): 平均(8/9) = 平均(0.88888889) = 0
//
//
// 注意:
//
//
// 给定矩阵中的整数范围为 [0, 255]。
// 矩阵的长和宽的范围均为 [1, 150]。
//
// Related Topics 数组 矩阵 👍 85 👎 0

package algorithm_600

func imageSmoother(img [][]int) (res [][]int) {
	// 8联通
	var dxy = [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	var m, n = len(img), len(img[0])
	res = make([][]int, len(img))
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var num, total = 1, img[i][j]
			for k := range dxy {
				var ny, nx = i + dxy[k][0], j + dxy[k][1]
				if ny < 0 || ny >= m || nx < 0 || nx >= n {
					continue
				}
				num++
				total += img[ny][nx]
			}
			res[i][j] = total / num
		}
	}
	return res
}

func imageSmoother1(img [][]int) (res [][]int) {
	res = make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			var r, n = img[i][j], 1
			// 同一层左
			if j-1 >= 0 {
				r += img[i][j-1]
				n++
			}
			// 同一层右
			if j+1 < len(img[i]) {
				r += img[i][j+1]
				n++
			}
			// 前一层
			if i-1 >= 0 {
				r += img[i-1][j]
				n++
				if j-1 >= 0 {
					r += img[i-1][j-1]
					n++
				}
				if j+1 < len(img[i]) {
					r += img[i-1][j+1]
					n++
				}
			}
			// 后一层
			if i+1 < len(img) {
				r += img[i+1][j]
				n++
				if j-1 >= 0 {
					r += img[i+1][j-1]
					n++
				}
				if j+1 < len(img[i]) {
					r += img[i+1][j+1]
					n++
				}
			}
			res[i] = append(res[i], r/n)
		}
	}
	return
}
