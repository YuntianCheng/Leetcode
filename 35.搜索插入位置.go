/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
package main

import "fmt"

// @lc code=start
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	index := (start + end) / 2
	for start <= end {
		if nums[index] > target {
			end = index - 1
		} else if nums[index] == target {
			return index
		} else {
			start = index + 1
		}
		index = (start + end) / 2
	}
	nums = append(append(nums[:start], target), nums[start:]...)
	return start
}

// @lc code=end
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}
