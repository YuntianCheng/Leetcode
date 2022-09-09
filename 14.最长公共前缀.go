package main

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	prefixBuilder := strings.Builder{}
	length := 201
	for _, str := range strs {
		if len(str) < length {
			length = len(str)
		}
	}
	for i := 0; i < length; i++ {
		count := 0
		value := strs[0][i]
		for _, str := range strs {
			if str[i] != value {
				return prefixBuilder.String()
			}
			count++
		}
		if count == len(strs) {
			prefixBuilder.WriteByte(value)
		}
	}
	return prefixBuilder.String()
}

// @lc code=end
