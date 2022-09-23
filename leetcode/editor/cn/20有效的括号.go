//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 3500 👎 0

package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, 0, len(s)/2)
	top := -1
	for _, block := range s {
		switch block {
		case '{':
			stack = append(stack, '}')
			top++
		case '[':
			stack = append(stack, ']')
			top++
		case '(':
			stack = append(stack, ')')
			top++
		default:
			if top < 0 {
				return false
			}
			p := stack[top]
			stack = stack[:top]
			top--
			if p != block {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(isValid("()"))
}
