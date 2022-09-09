/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */
package main

import "fmt"

// @lc code=start
func generate(numRows int) (result [][]int) {
	lineOne := []int{1}
	result = append(result, lineOne)
	lineOne = append(append([]int{0}, lineOne...), 0)
	for i := 2; i <= numRows; i++ {
		newLine := []int{}
		for j := 0; j < len(lineOne)-1; j++ {
			newLine = append(newLine, lineOne[j]+lineOne[j+1])
		}
		lineOne = append(append([]int{0}, newLine...), 0)
		result = append(result, newLine)
	}
	return
}

// @lc code=end
func main() {
	fmt.Println(generate(5))
}
