//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//
// 回文串不一定是字典当中的单词。
//
//
//
// 示例1：
//
// 输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
//
//
//
//
// Related Topics 位运算 哈希表 字符串 👍 92 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	var dict [128]int
	for i := range s {
		dict[s[i]]++
	}
	var count int
	for i := range dict {
		if dict[i]%2 != 0 {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
