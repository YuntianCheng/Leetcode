//给定一个整数数组
// prices，其中第 prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//
//
//输入: prices = [1]
//输出: 0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1323 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
//func maxProfit(prices []int) int {
//	dp := make([]int, len(prices))
//	for i := 1; i < len(prices); i++ {
//		max := dp[i-1]
//		for j := 1; j <= i; j++ {
//			trans := prices[i] - prices[i-j]
//			if trans < 0 {
//				continue
//			}
//			if i-j-2 >= 0 {
//				trans += dp[i-j-2]
//			}
//			if trans > max {
//				max = trans
//			}
//		}
//		dp[i] = max
//	}
//	return dp[len(prices)-1]
//}
func maxProfit232(prices []int) int {
	hold := make([]int, len(prices))
	unhold := make([]int, len(prices))
	hold[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		if i < 2 {
			hold[i] = hold[i-1]
			if hold[i] < -prices[i] {
				hold[i] = -prices[i]
			}
		} else {
			hold[i] = hold[i-1]
			if hold[i] < unhold[i-2]-prices[i] {
				hold[i] = unhold[i-2] - prices[i]
			}
		}
		unhold[i] = unhold[i-1]
		if unhold[i] < hold[i-1]+prices[i] {
			unhold[i] = hold[i-1] + prices[i]
		}
	}
	return unhold[len(prices)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
