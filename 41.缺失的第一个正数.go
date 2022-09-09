/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func firstMissingPositive(nums []int) int {
	min := 1
	max := len(nums)
	for i := 0; i < max; i++ {
		if nums[i] <= 0 {
			nums[i] = 1000
		}
	}
	for i := 0; i < max; i++ {
		q := math.Abs(float64(nums[i]))
		if int(q) <= max && nums[int(q)-1] > 0 {
			nums[int(q)-1] = -nums[int(q)-1]
		}
	}
	for {
		if min > max {
			break
		}
		if nums[min-1] < 0 {
			min++
		} else {
			break
		}
	}
	return min
}

// @lc code=end
func main() {
	fmt.Println(firstMissingPositive([]int{1, 1}))
}
