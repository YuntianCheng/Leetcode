/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */
package main

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	left := m
	right := n
	for _, val := range ops {
		if val[0] < left {
			left = val[0]
		}
		if right > val[1] {
			right = val[1]
		}
	}
	return left * right
}

// @lc code=end
