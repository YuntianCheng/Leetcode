/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
package main

import "fmt"

// @lc code=start
func generateMatrix(n int) [][]int {
	nums := 1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	top := 0
	left := 0
	bottom := n
	right := n
	for nums <= n*n {
		for i := left; i < right; i++ {
			result[top][i] = nums
			nums++
		}
		top++
		for i := top; i < bottom; i++ {
			result[i][right-1] = nums
			nums++
		}
		right--
		for i := right - 1; i >= left; i-- {
			result[bottom-1][i] = nums
			nums++
		}
		bottom--
		for i := bottom - 1; i >= top; i-- {
			result[i][left] = nums
			nums++
		}
		left++
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(generateMatrix(3))
}
