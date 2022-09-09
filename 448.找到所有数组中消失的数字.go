/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */
package main

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	st := make(map[int]bool, 0)
	for _, num := range nums {
		st[num] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := st[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end
