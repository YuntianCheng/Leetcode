/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */
package main

import "sort"

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)
	hNum := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > hNum {
			hNum++
		}
		if citations[i] == hNum {
			return citations[i]
		}
	}
	return hNum
}

// @lc code=end
