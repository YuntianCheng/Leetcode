//实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
//
// 示例 1：
//
//
//输入: s = "leetcode"
//输出: false
//
//
// 示例 2：
//
//
//输入: s = "abc"
//输出: true
//
//
// 限制：
//
//
// 0 <= len(s) <= 100
// s[i]仅包含小写字母
// 如果你不使用额外的数据结构，会很加分。
//
//
// Related Topics 位运算 哈希表 字符串 排序 👍 239 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func isUnique(astr string) bool {
	var have, cur = 0, 0
	for i := 0; i < len(astr); i++ {
		cur = 1 << (astr[i] - 'a')
		if have&cur > 0 {
			return false
		}
		have = have | cur
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
