//给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 本题中，将空字符串定义为有效的 回文串 。
//
//
//
// 示例 1:
//
//
//输入: s = "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
// 示例 2:
//
//
//输入: s = "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
//
//
//
//
// 注意：本题与主站 125 题相同： https://leetcode-cn.com/problems/valid-palindrome/
//
// Related Topics 双指针 字符串 👍 34 👎 0

package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome1(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	builder := strings.Builder{}
	for i := range s {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			builder.WriteByte(s[i])
		}
	}
	newS := builder.String()
	var i, j = 0, len(newS) - 1
	for i < j {
		if newS[i] != newS[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
