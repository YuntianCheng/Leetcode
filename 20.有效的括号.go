package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				if (left == '{' && v == '}') || (left == '[' && v == ']') || (left == '(' && v == ')') {
					stack = stack[0 : len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// @lc code=end
