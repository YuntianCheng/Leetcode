//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
// Related Topics 栈 字符串 动态规划 👍 2008 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	nums := make([]int, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			tmp := i - nums[i-1] - 1
			if tmp >= 0 && s[tmp] == '(' {
				nums[i] = nums[i-1] + 2
				if tmp > 0 {
					nums[i] += nums[tmp-1]
				}
			}
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(longestValidParentheses("(()(((()"))
}
