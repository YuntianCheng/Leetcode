/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */
package main

// @lc code=start
func minMoves(nums []int) int {
	min := nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}
	return sum - len(nums)*min
}

// @lc code=end
