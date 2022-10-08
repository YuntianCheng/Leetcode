//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 回溯 👍 1273 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func partition131(s string) [][]string {
	dp1 := make([][]bool, len(s))
	for i, _ := range dp1 {
		dp1[i] = make([]bool, len(s))
		dp1[i][i] = true
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i == 1 {
					dp1[i][j] = true
				} else {
					dp1[i][j] = dp1[i+1][j-1]
				}
			}
		}
	}
	dp := make([][][]string, len(s))
	dp[0] = [][]string{{string(s[0])}}
	for i := 1; i < len(s); i++ {
		tmp := append([][]string{}, dp[i-1]...)
		for j, _ := range tmp {
			tmp[j] = append(tmp[j], string(s[i]))
		}
		for j := 1; j < i; j++ {
			if dp1[j][i] {
				for _, v := range dp[j-1] {
					tmp = append(tmp, append(append([]string{}, v...), s[j:i+1]))
				}
			}
		}
		if dp1[0][i] {
			tmp = append(tmp, []string{s[0 : i+1]})
		}
		dp[i] = tmp
		fmt.Println(dp[i])
	}
	return dp[len(s)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(partition("ccaacabacb"))
}
