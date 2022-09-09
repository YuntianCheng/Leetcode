/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */
package main

import "fmt"

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, num := range nums {
		if num == 0 {
			if max < count {
				max = count
			}
			count = 0
		} else {
			count++
		}
	}
	if max < count {
		max = count
	}
	return max
}

// @lc code=end
func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
