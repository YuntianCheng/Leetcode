//给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
//
// 注意：此题 matrix 输入格式为一维 01 字符串数组。
//
//
//
// 示例 1：
//
//
//
//
//输入：matrix = ["10100","10111","11111","10010"]
//输出：6
//解释：最大矩形如上图所示。
//
//
// 示例 2：
//
//
//输入：matrix = []
//输出：0
//
//
// 示例 3：
//
//
//输入：matrix = ["0"]
//输出：0
//
//
// 示例 4：
//
//
//输入：matrix = ["1"]
//输出：1
//
//
// 示例 5：
//
//
//输入：matrix = ["00"]
//输出：0
//
//
//
//
// 提示：
//
//
// rows == matrix.length
// cols == matrix[0].length
// 0 <= row, cols <= 200
// matrix[i][j] 为 '0' 或 '1'
//
//
//
//
// 注意：本题与主站 85 题相同（输入参数格式不同）： https://leetcode-cn.com/problems/maximal-
//rectangle/
//
// Related Topics 栈 数组 动态规划 矩阵 单调栈 👍 62 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix []string) int {
	numTrix := make([][]int, len(matrix))
	for i := range numTrix {
		numTrix[i] = make([]int, len(matrix[i]))
	}
	for i := range numTrix {
		for j := range numTrix[i] {
			if i > 0 {
				if matrix[i][j] == '1' {
					numTrix[i][j] = numTrix[i-1][j] + 1
				}
			} else {
				if matrix[i][j] == '1' {
					numTrix[i][j] = 1
				}
			}
		}
	}
	var max1 int
	for _, v := range numTrix {
		var stack = make([]int, 0)
		var left = make([]int, len(v))
		var right = make([]int, len(v))
		for i, h := range v {
			for len(stack) > 0 && v[stack[len(stack)-1]] >= h {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				left[i] = -1
			} else {
				left[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		stack = stack[:0]
		for i := len(v) - 1; i >= 0; i-- {
			for len(stack) > 0 && v[stack[len(stack)-1]] >= v[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				right[i] = len(v)
			} else {
				right[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		for i := range left {
			if max1 < v[i]*(right[i]-left[i]-1) {
				max1 = v[i] * (right[i] - left[i] - 1)
			}
		}
	}
	return max1
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maximalRectangle([]string{"10100", "10111", "11111", "10010"}))
}
