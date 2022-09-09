/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func setZeroes(matrix [][]int) {
	zerosRow := ""
	zerosCol := ""
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zerosRow += strconv.Itoa(i) + ","
				zerosCol += strconv.Itoa(j) + ","
			}

		}
	}
	if len(zerosRow) > 0 {
		zerosRow = zerosRow[:len(zerosRow)-1]
		for _, val := range strings.Split(zerosRow, ",") {
			j, _ := strconv.Atoi(val)
			for k := 0; k < len(matrix[j]); k++ {
				matrix[j][k] = 0
			}
		}
	}
	if len(zerosCol) > 0 {
		zerosCol = zerosCol[:len(zerosCol)-1]
		for _, val := range strings.Split(zerosCol, ",") {
			j, _ := strconv.Atoi(val)
			for k := 0; k < len(matrix); k++ {
				matrix[k][j] = 0
			}
		}
	}
}

// @lc code=end
func main() {
	setZeroes([][]int{{1, 9, 7, 3, 8, 2, 1, 9, 8, 8, 2, 2, 9, 2}, {2, 4, 6, 7, 7, 5, 7, 8, 3, 5, 1, 5, 9, 0}, {9, 0, 9, 2, 1, 8, 6, 9, 1, 7, 1, 1, 7, 3}, {3, 6, 7, 7, 2, 3, 6, 9, 9, 4, 7, 4, 1, 0}, {0, 0, 0, 1, 3, 2, 4, 6, 0, 0, 6, 3, 7, 7}, {6, 6, 2, 0, 4, 9, 9, 7, 4, 6, 6, 5, 0, 6}})
}
