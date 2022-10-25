//给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 变位词 指字母相同，但排列不同的字符串。
//
//
//
// 示例 1：
//
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。
//
//
// 示例 2：
//
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的变位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的变位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的变位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
//
//
//
//
// 注意：本题与主站 438 题相同： https://leetcode-cn.com/problems/find-all-anagrams-in-a-
//string/
//
// Related Topics 哈希表 字符串 滑动窗口 👍 38 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	var result []int
	if len(s) < len(p) {
		return result
	}
	var sta [26]int
	for i := range p {
		sta[p[i]-'a']++
		sta[s[i]-'a']--
	}
	var diff int
	for _, v := range sta {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		result = append(result, 0)
	}
	for i := len(p); i < len(s); i++ {
		var out, in = s[i-len(p)], s[i]
		if out != in {
			if sta[out-'a'] == 0 {
				diff++
			}
			sta[out-'a']++
			if sta[out-'a'] == 0 {
				diff--
			}
			if sta[in-'a'] == 0 {
				diff++
			}
			sta[in-'a']--
			if sta[in-'a'] == 0 {
				diff--
			}
		}
		if diff == 0 {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
