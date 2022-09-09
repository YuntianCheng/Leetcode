/*
 * @lc app=leetcode.cn id=1156 lang=golang
 *
 * [1156] 单字符重复子串的最大长度
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func maxRepOpt1(text string) int {
	length := len(text)
	if length == 1 {
		return length
	}
	max := 0
	i := 0
	for {
		if i >= length-1 {
			break
		}
		num := 1
		j := i + 1
		current := text[i]
		for ; j < length; j++ {
			if text[j] == current {
				num++
			} else {
				break
			}
		}
		k := j + 1
		if k < length && strings.Contains(text[j:], string(current)) {
			for ; k < length; k++ {
				if text[k] == current {
					num++
				} else {
					break
				}
			}
		}
		if (k < length && strings.Contains(text[k:], string(current))) || strings.Contains(text[:i], string(current)) {
			num++
		}
		if max < num {
			max = num
		}
		i = j
	}
	return max
}

// @lc code=end
func main() {
	fmt.Println(maxRepOpt1("a"))
}
