package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	sort.Ints(nums)
	i := len(nums) - 1
	if len(nums) < 3 {
		return nums[i]
	}
	count := 0
	for {
		if i <= 0 || count == 3 {
			break
		}
		if nums[i] > nums[i-1] {
			count++
			if count == 2 {
				count++
			}
		}
		i--
	}
	if count == 3 {
		return nums[i]
	}
	return nums[len(nums)-1]
}

// @lc code=end
func main() {
	fmt.Println(thirdMax([]int{1, 2, 2, 5, 3, 5}))
}
