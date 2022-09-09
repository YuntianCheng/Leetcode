package main

import "strings"

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}
	numbers := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	collection := make([]string, 0)
	for _, b := range digits {
		collection = append(collection, numbers[b])
	}
	length := 1
	for _, c := range collection {
		length = length * len(c)
	}
	outputs := make([]string, length)
	cursor := make([]int, len(collection))
	for i := 0; i < length; i++ {
		output := strings.Builder{}
		for j := 0; j < len(cursor); j++ {
			output.WriteByte(collection[j][cursor[j]])
		}
		outputs[i] = output.String()
		for j := len(cursor) - 1; j >= 0; j-- {
			if cursor[j] < len(collection[j])-1 {
				cursor[j]++
				break
			}
			cursor[j] = 0
		}
	}
	return outputs
}

// @lc code=end
