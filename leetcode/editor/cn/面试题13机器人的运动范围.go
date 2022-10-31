//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一
//格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但
//它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
//
//
// 示例 1：
//
// 输入：m = 2, n = 3, k = 1
//输出：3
//
//
// 示例 2：
//
// 输入：m = 3, n = 1, k = 0
//输出：1
//
//
// 提示：
//
//
// 1 <= n,m <= 100
// 0 <= k <= 20
//
//
// Related Topics 深度优先搜索 广度优先搜索 动态规划 👍 575 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func bitSum(a int) int {
	return a%10 + (a/10)%10
}
func movingCount(m int, n int, k int) int {
	var dfsM func(int, int) int
	var visited = make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dfsM = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if visited[i][j] == true {
			return 0
		}
		visited[i][j] = true
		if bitSum(i)+bitSum(j) > k {
			return 0
		}
		return 1 + dfsM(i-1, j) + dfsM(i+1, j) + dfsM(i, j-1) + dfsM(i, j+1)
	}
	return dfsM(0, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
