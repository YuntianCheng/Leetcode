//二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
//
// 请返回 封闭岛屿 的数目。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,
//0,1],[1,1,1,1,1,1,1,0]]
//输出：2
//解释：
//灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
//
// 示例 2：
//
//
//
//
//输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
//输出：1
//
//
// 示例 3：
//
//
//输入：grid = [[1,1,1,1,1,1,1],
//             [1,0,0,0,0,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,1,0,1,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,0,0,0,0,1],
//             [1,1,1,1,1,1,1]]
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= grid.length, grid[0].length <= 100
// 0 <= grid[i][j] <=1
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 162 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func isClosedIsland(grid [][]int, i int, j int) bool {
	var left, right, top, bottom bool
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
		if grid[i][j] == 0 {
			grid[i][j] = -1
			left = isClosedIsland(grid, i, j-1)
			right = isClosedIsland(grid, i, j+1)
			top = isClosedIsland(grid, i-1, j)
			bottom = isClosedIsland(grid, i+1, j)
			if left && right && top && bottom {
				return true
			}
		} else if grid[i][j] == 1 || grid[i][j] == -1 {
			return true
		}
	}
	return false
}

func closedIsland(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 && isClosedIsland(grid, i, j) {
				count++
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}))
}
