//你现在手里有一份大小为
// n x n 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地。
//
// 请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的，并返回该距离。如果网格上只有陆地或者海洋，请返回 -1。
//
// 我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 -
//x1| + |y0 - y1| 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,0,1],[0,0,0],[1,0,1]]
//输出：2
//解释：
//海洋单元格 (1, 1) 和所有陆地单元格之间的距离都达到最大，最大距离为 2。
//
//
// 示例 2：
//
//
//
//
//输入：grid = [[1,0,0],[0,0,0],[0,0,0]]
//输出：4
//解释：
//海洋单元格 (2, 2) 和所有陆地单元格之间的距离都达到最大，最大距离为 4。
//
//
//
//
// 提示：
//
//
//
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 100
// grid[i][j] 不是 0 就是 1
//
//
// Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 295 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

type Point struct {
	X        int
	Y        int
	Distance int
}

func maxDistance(grid [][]int) int {
	queue := make([]Point, 0)
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 1 {
				queue = append(queue, Point{X: i, Y: j, Distance: 0})
			}
		}
	}
	start := 0
	end := len(queue) - 1
	xOffset := []int{-1, 1, 0, 0}
	yOffset := []int{0, 0, -1, 1}
	var distance int
	var hasWater bool
	for start <= end {
		x := queue[start].X
		y := queue[start].Y
		distance = queue[start].Distance
		start++
		for i := 0; i < 4; i++ {
			nextX := x + xOffset[i]
			nextY := y + yOffset[i]
			if nextY >= 0 && nextX >= 0 && nextX < len(grid) && nextY < len(grid[nextX]) && grid[nextX][nextY] == 0 {
				grid[nextX][nextY] = 1
				hasWater = true
				queue = append(queue, Point{nextX, nextY, distance + 1})
				end++
			}
		}
	}
	if hasWater {
		return distance
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
}
