//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 2618 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if j >= 1 {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
					if dp[i][j] > dp[i][j-1]+1 {
						dp[i][j] = dp[i][j-1] + 1
					}
					if dp[i][j] > dp[i-1][j] {
						dp[i][j] = dp[i-1][j] + 1
					}
				}
			} else {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minDistance("horse", "ros"))
}
