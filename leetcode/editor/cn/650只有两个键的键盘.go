//最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：
//
//
// Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
// Paste（粘贴）：粘贴 上一次 复制的字符。
//
//
// 给你一个数字 n ，你需要使用最少的操作次数，在记事本上输出 恰好 n 个 'A' 。返回能够打印出 n 个 'A' 的最少操作次数。
//
//
//
// 示例 1：
//
//
//输入：3
//输出：3
//解释：
//最初, 只有一个字符 'A'。
//第 1 步, 使用 Copy All 操作。
//第 2 步, 使用 Paste 操作来获得 'AA'。
//第 3 步, 使用 Paste 操作来获得 'AAA'。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 数学 动态规划 👍 482 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func minSteps(n int) int {
	a := make([]int, n)
	b := make([]int, 0)
	for i := 2; i <= n; i++ {
		var flag bool
		for j := 0; j < len(b); j++ {
			if i%b[j] == 0 {
				a[i-1] = a[i/b[j]-1] + b[j]
				flag = true
				break
			}
		}
		if !flag {
			a[i-1] = i
			b = append(b, i)
		}
	}
	return a[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
