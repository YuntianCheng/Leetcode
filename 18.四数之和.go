package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	results := new([][]int)
	for i := 0; i <= len(nums)-4; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		res := []int{nums[i]}
		sum(target-nums[i], 3, nums[i+1:], results, res)
	}
	return *results
}
func sum(target int, num int, subNums []int, results *[][]int, result []int) {
	if num == 1 {
		for _, n := range subNums {
			if n > target {
				return
			}
			if n == target {
				result = append(result, n)
				*results = append(*results, result)
				return
			}
		}
	} else {
		for i := 0; i <= len(subNums)-num; i++ {
			if i != 0 && subNums[i] == subNums[i-1] {
				continue
			}
			var temp []int
			temp = append(temp, result...)
			temp = append(temp, subNums[i])
			sum(target-subNums[i], num-1, subNums[i+1:], results, temp)
		}
	}
}

// @lc code=end
