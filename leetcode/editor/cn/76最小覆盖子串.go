//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//o(n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 2175 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	m := make(map[uint8]int, 0)
	for i, _ := range t {
		if _, ok := m[t[i]]; ok {
			m[t[i]]++
		} else {
			m[t[i]] = 1
		}
	}
	leftNum := len(t)
	var start, end = -1, -1
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if start == -1 {
				start = i
			}
			m[s[i]]--
			if m[s[i]] >= 0 {
				leftNum--
			}
			if leftNum == 0 {
				end = i
				break
			}
		}
	}
	if start == -1 || end == -1 {
		return ""
	}
	for j := start; j <= len(s)-len(t); j++ {
		if _, ok := m[s[j]]; ok {
			if m[s[j]] < 0 {
				m[s[j]]++
			} else {
				start = j
				break
			}
		}
	}
	result := s[start : end+1]
	for i := end + 1; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if s[i] == s[start] {
				end = i
				for j := start + 1; j <= len(s)-len(t); j++ {
					if _, ok = m[s[j]]; ok {
						if m[s[j]] == 0 {
							start = j
							if len(result) > end-start+1 {
								result = s[start : end+1]
							}
							break
						} else {
							m[s[j]]++
						}
					}
				}
			} else {
				m[s[i]]--
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minWindow("bba", "ab"))
}
