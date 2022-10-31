//硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
//
// 示例1:
//
//
// 输入: n = 5
// 输出：2
// 解释: 有两种方式可以凑成总金额:
//5=5
//5=1+1+1+1+1
//
//
// 示例2:
//
//
// 输入: n = 10
// 输出：4
// 解释: 有四种方式可以凑成总金额:
//10=10
//10=5+5
//10=5+1+1+1+1+1
//10=1+1+1+1+1+1+1+1+1+1
//
//
// 说明：
//
// 注意:
//
// 你可以假设：
//
//
// 0 <= n (总金额) <= 1000000
//
//
// Related Topics 数组 数学 动态规划 👍 261 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func waysToChange(n int) int {
	if n < 5 {
		return 1
	}
	var dp = make([]int, n+1)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}
	coins := []int{1, 5, 10, 25}
	for _, coin := range coins {
		for i := 5; i < n+1; i++ {
			if i-coin >= 0 {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[n] % 1000000007
}

//leetcode submit region end(Prohibit modification and deletion)
