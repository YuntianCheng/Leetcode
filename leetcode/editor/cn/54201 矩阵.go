//给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
//
// 两个相邻元素间的距离为 1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
//输出：[[0,0,0],[0,1,0],[0,0,0]]
//
//
// 示例 2：
//
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
//输出：[[0,0,0],[0,1,0],[1,2,1]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10⁴
// 1 <= m * n <= 10⁴
// mat[i][j] is either 0 or 1.
// mat 中至少有一个 0
//
//
// Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 756 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
type M struct {
	X int
	Y int
}

func updateMatrix(mat [][]int) [][]int {
	known := make([]M, 0)
	for i, _ := range mat {
		for j, _ := range mat[i] {
			if mat[i][j] == 0 {
				known = append(known, M{i, j})
			} else {
				mat[i][j] = len(mat) + len(mat[i])
			}
		}
	}
	xOffset := []int{-1, 1, 0, 0}
	yOffset := []int{0, 0, -1, 1}
	start := 0
	end := len(known)
	for start < end {
		x := known[start].X
		y := known[start].Y
		start++
		for i := 0; i < 4; i++ {
			nextX := x + xOffset[i]
			nextY := y + yOffset[i]
			if nextX >= 0 && nextX < len(mat) && nextY >= 0 && nextY < len(mat[nextX]) && mat[nextX][nextY] > mat[x][y]+1 {
				known = append(known, M{nextX, nextY})
				mat[nextX][nextY] = mat[x][y] + 1
				end++
			}
		}
	}
	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}
