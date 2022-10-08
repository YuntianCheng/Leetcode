//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//
//
//
// 示例 1：
//
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：0
//
//
// 示例 3：
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
// 0 <= n <= 5 * 10⁶
//
//
// Related Topics 数组 数学 枚举 数论 👍 957 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	primes := make([]int, 1)
	primes[0] = 2
	count := 1
	for i := 3; i < n; i++ {
		var b bool
		for j := 0; primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				b = true
				break
			}
		}
		if !b {
			count++
			primes = append(primes, i)
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(countPrimes(10))
}
