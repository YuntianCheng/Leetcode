//给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。
//
// 注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。
//
//
//
// 示例 1:
//
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
//
//
//输入: strs = [""]
//输出: [[""]]
//
//
// 示例 3:
//
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
//
//
//
//
// 注意：本题与主站 49 题相同： https://leetcode-cn.com/problems/group-anagrams/
//
// Related Topics 数组 哈希表 字符串 排序 👍 40 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	var result = make([][]string, 0)
	var m = make(map[[26]int][]string, 0)
	for _, str := range strs {
		var key [26]int
		for i := range str {
			key[str[i]-'a']++
		}
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
