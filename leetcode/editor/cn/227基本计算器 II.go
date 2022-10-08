//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 整数除法仅保留整数部分。
//
// 你可以假设给定的表达式总是有效的。所有中间结果将在 [-2³¹, 2³¹ - 1] 的范围内。
//
// 注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
//
//
//
// 示例 1：
//
//
//输入：s = "3+2*2"
//输出：7
//
//
// 示例 2：
//
//
//输入：s = " 3/2 "
//输出：1
//
//
// 示例 3：
//
//
//输入：s = " 3+5 / 2 "
//输出：5
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
// s 表示一个 有效表达式
// 表达式中的所有整数都是非负整数，且在范围 [0, 2³¹ - 1] 内
// 题目数据保证答案是一个 32-bit 整数
//
//
// Related Topics 栈 数学 字符串 👍 622 👎 0

package main

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	suffix := make([]int, 0)
	stack := make([]uint8, 0)
	top1 := -1
	top2 := -1
	num := 0
	s = strings.Trim(s, " ")
	for i, _ := range s {
		if s[i] >= '0' && s[i] <= '9' {
			n, _ := strconv.Atoi(string(s[i]))
			num = 10*num + n
			if i == len(s)-1 {
				if top2 < len(suffix)-1 {
					top2++
					suffix[top2] = num
				} else {
					top2++
					suffix = append(suffix, num)
				}
				num = 0
			}
		} else {
			if s[i] != ' ' {
				if top2 < len(suffix)-1 {
					top2++
					suffix[top2] = num
				} else {
					top2++
					suffix = append(suffix, num)
				}
				num = 0
				if top1 >= 0 {
					switch s[i] {
					case '+', '-':
						for top1 >= 0 {
							b := suffix[top2]
							top2--
							a := suffix[top2]
							top2--
							var r int
							switch stack[top1] {
							case '+':
								r = a + b
							case '-':
								r = a - b
							case '*':
								r = a * b
							case '/':
								r = a / b
							}
							top1--
							top2++
							suffix[top2] = r
						}
					case '*', '/':
						if stack[top1] == '*' || stack[top1] == '/' {
							b := suffix[top2]
							top2--
							a := suffix[top2]
							top2--
							var r int
							switch stack[top1] {
							case '*':
								r = a * b
							case '/':
								r = a / b
							}
							top1--
							top2++
							suffix[top2] = r
						}
					}
				}
				if top1 < len(stack)-1 {
					top1++
					stack[top1] = s[i]
				} else {
					top1++
					stack = append(stack, s[i])
				}
			}
		}
	}
	var r int
	for top1 >= 0 {
		b := suffix[top2]
		top2--
		a := suffix[top2]
		top2--
		switch stack[top1] {
		case '+':
			r = a + b
		case '-':
			r = a - b
		case '*':
			r = a * b
		case '/':
			r = a / b
		}
		top2++
		top1--
		suffix[top2] = r
	}
	return suffix[0]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(calculate("2*3+4"))
}
