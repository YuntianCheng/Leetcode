//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//
// 返回符合要求的 最少分割次数 。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：0
//
//
// 示例 3：
//
//
//输入：s = "ab"
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 621 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	dp1 := make([][]bool, len(s))
	for i, _ := range dp1 {
		dp1[i] = make([]bool, len(s))
		dp1[i][i] = true
	}
	//dp2 := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if i+1 < j {
					dp1[i][j] = dp1[i+1][j-1]
				} else {
					dp1[i][j] = true
				}
			}
		}
	}
	dp2 := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if !dp1[0][i] {
			dp2[i] = len(s)
			for j := 1; j <= i; j++ {
				if dp1[j][i] && dp2[j-1]+1 < dp2[i] {
					dp2[i] = dp2[j-1] + 1
				}
			}
		}
	}
	return dp2[len(s)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minCut("aab"))
}
