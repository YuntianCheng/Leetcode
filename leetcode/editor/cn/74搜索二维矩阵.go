//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
//
//
// Related Topics 数组 二分查找 矩阵 👍 719 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)

func bS(nums []int, target, start, end int) bool {
	if start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if target < nums[mid] {
			return bS(nums, target, start, mid-1)
		} else {
			return bS(nums, target, mid+1, end)
		}
	}
	if nums[start] == target {
		return true
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if target <= matrix[i][len(matrix[i])-1] && target >= matrix[i][0] {
			return bS(matrix[i], target, 0, len(matrix[i])-1)
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
