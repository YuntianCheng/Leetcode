package main

import "strings"

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	i := 0
	j := 0
	for {
		if i >= len(s) {
			break
		}
		if j >= len(p) {
			break
		}
		if p[j] != '.' && p[j] != '*' {
			if p[j] != s[i] {
				if j < len(p)-1 && p[j+1] == '*' {
					j++
				} else {
					return false
				}
			}
			i++
			j++
			continue
		}
		if p[j] == '.' {
			j++
			i++
			continue
		}
		if p[j] == '*' {
			if p[j-1] == '.' && j == len(p)-1 {
				i = len(s)
				j = len(p)
				break
			}
			char := s[i-1]
			for {
				if i >= len(s) {
					break
				}
				if j+1 < len(p) && p[j+1] == '*' {
					j++
					continue
				}
				if s[i] == char {
					i++
					continue
				}
				break
			}
			j++
		}
	}
	if i <= len(s)-1 {
		return false
	}
	if j <= len(p)-1 && len(strings.Replace(s[j:], "*", "", -1)) > 0 {
		return false
	}
	return true
}

// @lc code=end
