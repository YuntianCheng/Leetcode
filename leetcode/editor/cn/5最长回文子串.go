//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 5695 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome1(s string) string {
	var maxOu string
	var maxJi string
	for i := 0; i < len(s); i++ {
		ouTmpLen := 0
		jiTmpLen := 0
		for j := 0; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] == s[i+j] {
				jiTmpLen = 2*j + 1
				if jiTmpLen > len(maxJi) {
					maxJi = s[i-j : i+j+1]
				}
			} else {
				break
			}
		}
		for j := 0; i-j >= 0 && i+j+1 < len(s); j++ {
			if s[i-j] == s[i+j+1] {
				ouTmpLen = 2*j + 2
				if ouTmpLen > len(maxOu) {
					maxOu = s[i-j : i+j+2]
				}
			} else {
				break
			}
		}
	}
	if len(maxOu) > len(maxJi) {
		return maxOu
	}
	return maxJi
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
