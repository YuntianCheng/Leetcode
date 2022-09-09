/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, 0)
	for index, value := range nums {
		if _, ok := indexMap[target-value]; ok {
			return []int{index, indexMap[target-value]}
		}
		if _, ok := indexMap[value]; !ok {
			indexMap[value] = index
		}
	}
	return []int{}
}

// @lc code=end
