//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字符串 👍 2482 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	length := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < length {
			length = len(strs[i])
		}
	}
	result := ""
	for i := 0; i < length; i++ {
		char := strs[0][i]
		common := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				common = false
				break
			}
		}
		if common {
			result = result + string(char)
		} else {
			break
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
