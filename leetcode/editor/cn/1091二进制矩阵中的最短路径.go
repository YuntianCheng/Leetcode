//给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
//
// 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求
//：
//
//
// 路径途经的所有单元格都的值都是 0 。
// 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
//
//
// 畅通路径的长度 是该路径途经的单元格总数。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,1],[1,0]]
//输出：2
//
//
// 示例 2：
//
//
//输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
//输出：4
//
//
// 示例 3：
//
//
//输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
//输出：-1
//
//
//
//
// 提示：
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 100
// grid[i][j] 为 0 或 1
//
//
// Related Topics 广度优先搜索 数组 矩阵 👍 218 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

type Place struct {
	X int
	Y int
	D int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	visitedPath := make([]Place, 1)
	visitedPath[0] = Place{0, 0, 1}
	grid[0][0] = 1
	xOffset := []int{-1, 1, 0, 0, -1, -1, 1, 1}
	yOffset := []int{0, 0, -1, 1, -1, 1, -1, 1}
	start := 0
	end := len(visitedPath)
	for start < end {
		x := visitedPath[start].X
		y := visitedPath[start].Y
		d := visitedPath[start].D
		if x == len(grid)-1 && y == len(grid[x])-1 {
			return d
		}
		start++
		for i := 0; i < 8; i++ {
			nextX := x + xOffset[i]
			nextY := y + yOffset[i]
			if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[nextX]) && grid[nextX][nextY] == 0 {
				visitedPath = append(visitedPath, Place{nextX, nextY, d + 1})
				grid[nextX][nextY] = 1
				end++
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
}
