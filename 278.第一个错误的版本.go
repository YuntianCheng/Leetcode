/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */
package main

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start := 1
	end := n
	num := (start + end) / 2
	for start < end {
		if isBadVersion(num) {
			if !isBadVersion(num - 1) {
				return num
			}
			end = num - 1
			num = (start + end) / 2
		} else {
			start = num + 1
			num = (start + end) / 2
		}
	}
	return num
}

// @lc code=end
