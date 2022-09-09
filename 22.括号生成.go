package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	result := new([]string)
	K(0, n, result, "")
	return *result
}

func K(leftcount int, nums int, result *[]string, value string) {
	if leftcount == 0 && nums == 0 {
		*result = append(*result, value)
		return
	}
	if nums > 0 {
		K(leftcount+1, nums-1, result, value+"(")
	}
	if leftcount > 0 {
		K(leftcount-1, nums, result, value+")")
	}
}

// @lc code=end
func main() {
	fmt.Println(generateParenthesis(3))
}
