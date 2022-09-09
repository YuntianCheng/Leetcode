/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */
package main

// @lc code=start
import "math"

func maxSubArray(nums []int) int {
	sum := 0
	result := nums[0]
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		result = int(math.Max(float64(result), float64(sum)))
	}
	return result
}

// @lc code=end
