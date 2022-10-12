//设想有个机器人坐在一个网格的左上角，网格 r 行 c 列。机器人只能向下或向右移动，但不能走到一些被禁止的网格（有障碍物）。设计一种算法，寻找机器人从左上角
//移动到右下角的路径。
//
//
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
// 返回一条可行的路径，路径由经过的网格的行号和列号组成。左上角为 0 行 0 列。如果没有可行的路径，返回空数组。
//
// 示例 1:
//
// 输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: [[0,0],[0,1],[0,2],[1,2],[2,2]]
//解释:
//输入中标粗的位置即为输出表示的路径，即
//0行0列（左上角） -> 0行1列 -> 0行2列 -> 1行2列 -> 2行2列（右下角）
//
// 说明：r 和 c 的值均不超过 100。
//
// Related Topics 数组 动态规划 回溯 矩阵 👍 108 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func dfs0802(obstacleGrid [][]int, i, j int, path *[][]int, m *[][]bool) bool {
	if i+1 == len(obstacleGrid) && j+1 == len(obstacleGrid[i]) && obstacleGrid[i][j] == 0 {
		return true
	}
	if (*m)[i][j] {
		return false
	}
	if obstacleGrid[i][j] == 1 {
		(*m)[i][j] = true
		return false
	}
	if i+1 < len(obstacleGrid) {
		*path = append(*path, []int{i + 1, j})
		if dfs0802(obstacleGrid, i+1, j, path, m) {
			return true
		}
		(*m)[i+1][j] = true
		*path = (*path)[:len(*path)-1]
	}
	if j+1 < len(obstacleGrid[i]) {
		*path = append(*path, []int{i, j + 1})
		if dfs0802(obstacleGrid, i, j+1, path, m) {
			return true
		}
		(*m)[i][j+1] = true
		*path = (*path)[:len(*path)-1]
	}
	return false
}

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	if obstacleGrid[0][0] == 1 {
		return [][]int{}
	}
	path := make([][]int, 0)
	path = append(path, []int{0, 0})
	m := make([][]bool, len(obstacleGrid))
	for i := range m {
		m[i] = make([]bool, len(obstacleGrid[i]))
	}
	if dfs0802(obstacleGrid, 0, 0, &path, &m) {
		return path
	}
	return [][]int{}
}

//leetcode submit region end(Prohibit modification and deletion)
