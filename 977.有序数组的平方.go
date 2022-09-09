/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
package main

import (
	"fmt"
)

// @lc code=start
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	result := make([]int, len(nums))
	k := len(nums) - 1
	i := 0
	j := k
	for i <= j {
		if nums[i] < nums[j] {
			result[k] = nums[j]
			j--
		} else {
			result[k] = nums[i]
			i++
		}
		k--
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
