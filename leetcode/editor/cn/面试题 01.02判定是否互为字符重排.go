//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
// 示例 1：
//
// 输入: s1 = "abc", s2 = "bca"
//输出: true
//
//
// 示例 2：
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
//
//
// 说明：
//
//
// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100
//
//
// Related Topics 哈希表 字符串 排序 👍 140 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	var m [128]int
	for i := range s1 {
		m[s1[i]]++
	}
	for i := range s2 {
		if m[s2[i]] <= 0 {
			return false
		}
		m[s2[i]]--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
