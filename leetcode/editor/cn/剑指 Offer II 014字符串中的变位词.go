//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
//
//
//
// 示例 1：
//
//
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
//
//
// 示例 2：
//
//
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
//
//
//
//
//
// 注意：本题与主站 567 题相同： https://leetcode-cn.com/problems/permutation-in-string/
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 70 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	if s1 == s2 {
		return true
	}
	var c [26]int
	for i := 0; i < len(s1); i++ {
		c[s1[i]-'a']++
		c[s2[i]-'a']--
	}
	var diff int
	for _, v := range c {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		var out, in = s2[i-len(s1)], s2[i]
		if out == in {
			continue
		}
		if c[out-'a'] == 0 {
			diff++
		}
		c[out-'a']++
		if c[out-'a'] == 0 {
			diff--
		}
		if c[in-'a'] == 0 {
			diff++
		}
		c[in-'a']--
		if c[in-'a'] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
