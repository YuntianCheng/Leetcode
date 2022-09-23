//有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
//
// 这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上
//单元格 高于海平面的高度 。
//
// 岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
//
// 返回网格坐标 result 的 2D 列表 ，其中 result[i] = [ri, ci] 表示雨水从单元格 (ri, ci) 流动 既可流向太平洋也可
//流向大西洋 。
//
//
//
// 示例 1：
//
//
//
//
//输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
//输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//
//
// 示例 2：
//
//
//输入: heights = [[2,1],[1,2]]
//输出: [[0,0],[0,1],[1,0],[1,1]]
//
//
//
//
// 提示：
//
//
// m == heights.length
// n == heights[r].length
// 1 <= m, n <= 200
// 0 <= heights[r][c] <= 10⁵
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 👍 520 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

type LandPoint struct {
	X int
	Y int
}

func pacificAtlantic(heights [][]int) [][]int {
	result := make([][]int, 0)
	pQueue := make([]LandPoint, 0)
	aQueue := make([]LandPoint, 0)
	for i, _ := range heights {
		for j, _ := range heights[i] {
			heights[i][j]++
		}
	}
	for j := 0; j < len(heights[0]); j++ {
		pQueue = append(pQueue, LandPoint{0, j})
		heights[0][j] = -heights[0][j]
	}
	for i := 1; i < len(heights); i++ {
		pQueue = append(pQueue, LandPoint{i, 0})
		heights[i][0] = -heights[i][0]
	}

	xOffset := []int{1, 0, -1, 0}
	yOffset := []int{0, 1, 0, -1}
	start := 0
	end := len(pQueue)
	for start < end {
		x := pQueue[start].X
		y := pQueue[start].Y
		start++
		height := -heights[x][y]
		for i := 0; i < 4; i++ {
			nextX := x + xOffset[i]
			nextY := y + yOffset[i]
			if nextX < len(heights) && nextX >= 0 && nextY < len(heights[nextX]) && nextY >= 0 && heights[nextX][nextY] >= height {
				pQueue = append(pQueue, LandPoint{nextX, nextY})
				heights[nextX][nextY] = -heights[nextX][nextY]
				end++
			}
		}
	}
	for j := 0; j < len(heights[len(heights)-1]); j++ {
		if heights[len(heights)-1][j] < 0 {
			result = append(result, []int{len(heights) - 1, j})
			heights[len(heights)-1][j] = -heights[len(heights)-1][j] + 100001
		} else {
			heights[len(heights)-1][j] += 100001
		}
		aQueue = append(aQueue, LandPoint{len(heights) - 1, j})
	}
	for i := 0; i < len(heights)-1; i++ {
		if heights[i][len(heights[i])-1] < 0 {
			result = append(result, []int{i, len(heights[i]) - 1})
			heights[i][len(heights[i])-1] = -heights[i][len(heights[i])-1] + 100001
		} else {
			heights[i][len(heights[i])-1] += 100001
		}
		aQueue = append(aQueue, LandPoint{i, len(heights[i]) - 1})
	}
	start = 0
	end = len(aQueue)
	for start < end {
		x := aQueue[start].X
		y := aQueue[start].Y
		start++
		height := heights[x][y] - 100001
		for i := 0; i < 4; i++ {
			nextX := x + xOffset[i]
			nextY := y + yOffset[i]
			if nextX < len(heights) && nextX >= 0 && nextY < len(heights[nextX]) && nextY >= 0 {
				var nextHeight int
				if heights[nextX][nextY] < 0 {
					nextHeight = -heights[nextX][nextY]
				} else {
					nextHeight = heights[nextX][nextY]
				}
				if nextHeight <= 100001 && nextHeight >= height {
					aQueue = append(aQueue, LandPoint{nextX, nextY})
					end++
					if heights[nextX][nextY] < 0 {
						result = append(result, []int{nextX, nextY})
					}
					heights[nextX][nextY] = nextHeight + 100001
				}
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}))
}
