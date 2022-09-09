package main

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	start := 0
	end := len(nums) - 1
	mid := 1
	var bestAbs float64 = 100000
	var tempSum int
	sum := nums[start] + nums[end] + nums[mid]
	for start := 0; start <= len(nums)-3; start++ {
		var abs float64 = 10000
		end := len(nums) - 1
		for mid = start + 1; mid < end; {
			two := nums[start] + nums[end] - target
			tempAbs := math.Abs(float64(two + nums[mid]))
			if tempAbs <= abs {
				abs = tempAbs
				tempSum = nums[start] + nums[end] + nums[mid]
			}
			if abs == 0 {
				return tempSum
			}
			if nums[start]+nums[end]+nums[mid] >= target {
				end--
			} else {
				mid++
			}
		}
		if abs < bestAbs {
			bestAbs = abs
			sum = tempSum
		}
	}
	return sum
}

// @lc code=end

func main() {
	closest := threeSumClosest([]int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33}, 0)
	fmt.Println(closest)
}
