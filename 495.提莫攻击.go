/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */
package main

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	totalTime := 0
	length := len(timeSeries)
	for i := 0; i < length; i++ {
		stopTime := timeSeries[i] + duration - 1
		if i == length-1 {
			totalTime += duration
			break
		}
		if stopTime < timeSeries[i+1] {
			totalTime += duration
		} else {
			totalTime += timeSeries[i+1] - timeSeries[i]
		}
	}
	return totalTime
}

// @lc code=end
