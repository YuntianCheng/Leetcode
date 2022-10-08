//给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,4,7,5,3,6,8,9]
//
//
// 示例 2：
//
//
//输入：mat = [[1,2],[3,4]]
//输出：[1,2,3,4]
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
// -10⁵ <= mat[i][j] <= 10⁵
//
//
// Related Topics 数组 矩阵 模拟 👍 404 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findDiagonalOrder(mat [][]int) []int {
	var result []int
	n := len(mat) * len(mat[0])
	count := 0
	var up = true
	var i, j int
	for count < n {
		result = append(result, mat[i][j])
		count++
		if up {
			if i-1 >= 0 && j+1 < len(mat[i]) {
				i--
				j++
			} else {
				up = !up
				if j+1 < len(mat[i]) {
					j++
				} else {
					i++
				}
			}
		} else {
			if i+1 < len(mat) && j-1 >= 0 {
				i++
				j--
			} else {
				up = !up
				if i+1 < len(mat) {
					i++
				} else {
					j++
				}
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
