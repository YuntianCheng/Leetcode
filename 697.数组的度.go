/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */
package main

import "fmt"

// @lc code=start
func findShortestSubArray(nums []int) int {
	s := make(map[int]int, 0)
	for _, num := range nums {
		if _, ok := s[num]; ok {
			s[num]++
			continue
		}
		s[num] = 1
	}
	maxNum := 0
	maxValue := make([]int, 0)
	for _, value := range s {
		if value > maxNum {
			maxNum = value
		}
	}
	for key := range s {
		if s[key] == maxNum {
			maxValue = append(maxValue, key)
		}
	}
	length := 50000
	for _, v := range maxValue {
		i := 0
		j := len(nums) - 1
		for {
			if nums[i] != v {
				i++
			}
			if nums[j] != v {
				j--
			}
			if nums[i] == v && nums[j] == v {
				break
			}
		}
		if length > j-i+1 {
			length = j - i + 1
		}
	}
	return length
}

// @lc code=end
func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
}
