//给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
//
//
//
// 注意：本题与主站 3 题相同： https://leetcode-cn.com/problems/longest-substring-without-
//repeating-characters/
//
// Related Topics 哈希表 字符串 滑动窗口 👍 59 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	m := map[byte]struct{}{}
	var max int
	var left int
	for right := range s {
		if _, ok := m[s[right]]; ok {
			if max < right-left {
				max = right - left
			}
			for s[left] != s[right] {
				delete(m, s[left])
				left++
			}
			left++
		} else {
			m[s[right]] = struct{}{}
		}
	}
	if max < len(s)-left {
		max = len(s) - left
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
