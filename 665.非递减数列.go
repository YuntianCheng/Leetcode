/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */
package main

import "fmt"

// @lc code=start
func checkPossibility(nums []int) bool {
	exchanged := false
	var firstArray []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if exchanged {
				return false
			}
			exchanged = true
			firstArray = nums[:i]
		}
	}
	count := 0
	if len(firstArray) == len(nums)-1 {
		return true
	}
	for _, val := range firstArray {
		if val > nums[len(firstArray)] {
			count++
		}
	}
	if count > 1 && nums[len(firstArray)+1] <= firstArray[len(firstArray)-1] {
		return false
	}
	return true
}

// @lc code=end
func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}
