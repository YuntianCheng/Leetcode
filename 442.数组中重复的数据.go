/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		n := math.Abs(float64(num))
		if nums[int(n)-1] > 0 {
			nums[int(n)-1] *= -1
		} else {
			result = append(result, int(n))
		}
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(findDuplicates([]int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7}))
}
