package main

import "fmt"

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	result := new([]int)
	printM(matrix, result)
	return *result
}

func printM(matrix [][]int, result *[]int) {
	if len(matrix) > 0 {
		startRow := 0
		endCol := len(matrix[startRow]) - 1
		endRow := len(matrix) - 1
		startCol := 0
		*result = append(*result, matrix[startRow]...)
		if endRow <= startRow || endCol < startCol {
			return
		}
		for index, n := range matrix {
			if index != startRow {
				*result = append(*result, n[endCol])
			}
		}
		if endCol == startCol {
			return
		}
		for i := len(matrix[endRow]) - 2; i >= 0; i-- {
			*result = append(*result, matrix[endRow][i])
		}
		for i := len(matrix) - 2; i > startRow; i-- {
			*result = append(*result, matrix[i][startCol])
		}
		matrix = matrix[1 : len(matrix)-1]
		for index, _ := range matrix {
			matrix[index] = matrix[index][1 : len(matrix[index])-1]
		}
		printM(matrix, result)
	}
	return
}

// @lc code=end
func main() {
	closest := spiralOrder([][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}, {6, 16}, {7, 17}, {8, 18}, {9, 19}, {10, 20}})
	fmt.Println(closest)
}
