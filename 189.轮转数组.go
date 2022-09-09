/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */
package main

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, len(nums[:k])-1; i < j; i, j = i+1, j-1 {
		nums[:k][i], nums[:k][j] = nums[:k][j], nums[:k][i]
	}
	for i, j := 0, len(nums[k:])-1; i < j; i, j = i+1, j-1 {
		nums[k:][i], nums[k:][j] = nums[k:][j], nums[k:][i]
	}
}

// @lc code=end
func main() {
	rotate([]int{-1, -100, 3, 99}, 2)
}
