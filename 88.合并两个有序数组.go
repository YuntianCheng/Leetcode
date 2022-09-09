/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
package main

import "sort"

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m
	for _, v := range nums2 {
		nums1[i] = v
		i++
	}
	sort.Ints(nums1)
}

// @lc code=end
