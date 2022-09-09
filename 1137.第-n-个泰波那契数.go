/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */
package main

// @lc code=start
func tribonacci(n int) int {
	fibs := make([]int, 3)
	fibs[0] = 0
	fibs[1] = 1
	fibs[2] = 1
	for i := 3; i <= n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2]+fibs[i-3])
	}
	return fibs[n]
}

// @lc code=end
