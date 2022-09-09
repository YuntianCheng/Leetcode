/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */
package main

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	var result [][]int
	for row, rVal := range img {
		var r []int
		for col := range rVal {
			r = append(r, getAvg(row, col, img))
		}
		result = append(result, r)
	}
	return result
}
func getAvg(row int, col int, img [][]int) int {
	sum := 0
	nums := 0
	rowLength := len(img)
	colLength := len(img[0])
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < rowLength && j >= 0 && j < colLength {
				sum += img[i][j]
				nums++
			}
		}
	}
	return sum / nums
}

// @lc code=end
