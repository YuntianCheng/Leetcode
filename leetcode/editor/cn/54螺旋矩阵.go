//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 矩阵 模拟 👍 1222 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	result := new([]int)
	printM(matrix, result)
	return *result
}

func printM(matrix [][]int, result *[]int) {
	if len(matrix) > 0 {
		startRow := 0
		endCol := len(matrix[startRow]) - 1
		endRow := len(matrix) - 1
		startCol := 0
		*result = append(*result, matrix[startRow]...)
		if endRow <= startRow || endCol < startCol {
			return
		}
		for index, n := range matrix {
			if index != startRow {
				*result = append(*result, n[endCol])
			}
		}
		if endCol == startCol {
			return
		}
		for i := len(matrix[endRow]) - 2; i >= 0; i-- {
			*result = append(*result, matrix[endRow][i])
		}
		for i := len(matrix) - 2; i > startRow; i-- {
			*result = append(*result, matrix[i][startCol])
		}
		matrix = matrix[1 : len(matrix)-1]
		for index, _ := range matrix {
			matrix[index] = matrix[index][1 : len(matrix[index])-1]
		}
		printM(matrix, result)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
