/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */
package main

// @lc code=start
func fib(n int) int {
	fibs := make([]int, 2)
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i <= n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}
	return fibs[n]
}

// @lc code=end
