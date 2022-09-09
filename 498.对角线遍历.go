/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */
package main

// @lc code=start
func findDiagonalOrder(mat [][]int) (result []int) {
	n := len(mat) * len(mat[0])
	nums := 0
	top := true
	i := 0
	j := 0
	for nums < n {
		result = append(result, mat[i][j])
		nums++
		if top {
			if i-1 >= 0 && j+1 < len(mat[i]) {
				i--
				j++
			} else {
				top = !top
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
				top = !top
				if i+1 < len(mat) {
					i++
				} else {
					j++
				}
			}
		}
	}
	return
}

// @lc code=end
