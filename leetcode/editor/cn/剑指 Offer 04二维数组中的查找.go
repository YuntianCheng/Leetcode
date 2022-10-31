//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个
//整数，判断数组中是否含有该整数。
//
//
//
// 示例:
//
// 现有矩阵 matrix 如下：
//
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//
//
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false。
//
//
//
// 限制：
//
// 0 <= n <= 1000
//
// 0 <= m <= 1000
//
//
//
// 注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
//
// Related Topics 数组 二分查找 分治 矩阵 👍 812 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 || matrix[i][0] > target {
			break
		}
		var m, n = 0, len(matrix[i]) - 1
		for m <= n {
			mid := (m + n) / 2
			if matrix[i][mid] == target {
				return true
			}
			if matrix[i][mid] < target {
				m = mid + 1
			} else {
				n = mid - 1
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
