/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */
package main

import "fmt"

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		index := (start + end) / 2
		if (index == 0 || arr[index] > arr[index-1]) && (index == len(arr)-1 || arr[index] > arr[index+1]) {
			return index
		} else if arr[index] < arr[index+1] {
			start = index + 1
		} else if arr[index] < arr[index-1] {
			end = index - 1
		}
	}
	return 0
}

// @lc code=end
func main() {
	fmt.Println(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
}
