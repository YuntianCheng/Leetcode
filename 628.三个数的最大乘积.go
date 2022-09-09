/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */
package main

import "sort"

// @lc code=start
func maximumProduct(nums []int) int {
	length := len(nums)
	if length == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	sort.Ints(nums)
	max1 := nums[length-1] * nums[length-2] * nums[length-3]
	max2 := nums[length-1] * nums[0] * nums[1]
	if max1 > max2 {
		return max1
	}
	return max2

}

// @lc code=end
