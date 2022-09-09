/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
func moveZeroes(nums []int) {
	length := len(nums)
	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for {
		if j == length {
			break
		}
		nums[j] = 0
		j++
	}
}

// @lc code=end
