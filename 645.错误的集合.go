/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */
package main

import "fmt"

// @lc code=start
func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	statistics := make(map[int]int, 0)
	for index := range nums {
		statistics[index+1] = 0
	}
	for _, num := range nums {
		statistics[num]++
	}
	for index := range nums {
		if statistics[index+1] == 0 {
			result[1] = index + 1
		}
		if statistics[index+1] > 1 {
			result[0] = index + 1
		}
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(findErrorNums([]int{1, 1}))
}
