/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
package main

func s(nums []int, target int, start int, end int) int {
	if start < end {
		index := (end + start) / 2
		if nums[index] > target {
			return s(nums, target, start, index-1)
		} else if nums[index] == target {
			return index
		} else {
			return s(nums, target, index+1, end)
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}

func search(nums []int, target int) int {
	return s(nums, target, 0, len(nums)-1)
}

// @lc code=end
