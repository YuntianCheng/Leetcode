/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */
package main

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	start := 1
	end := n
	num := (start + end) / 2
	for start < end {
		r := guess(num)
		if r == 0 {
			return num
		} else if r == -1 {
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
