//给你一个整数 n ，请你找出并返回第 n 个 丑数 。
//
// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
//
//
//
// 示例 1：
//
//
//输入：n = 10
//输出：12
//解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//解释：1 通常被视为丑数。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1690
//
//
// Related Topics 哈希表 数学 动态规划 堆（优先队列） 👍 970 👎 0

package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	points := []int{0, 0, 0}
	dp[0] = 1
	i := 1
	for i < n {
		var p1, p2, p3 = dp[points[0]] * 2, dp[points[1]] * 3, dp[points[2]] * 5
		value := p1
		if p2 < value {
			value = p2
		}
		if p3 < value {
			value = p3
		}
		if p1 == value {
			points[0]++
		}
		if p2 == value {
			points[1]++
		}
		if p3 == value {
			points[2]++
		}
		dp[i] = value
		i++
	}
	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(nthUglyNumber(37))
}
