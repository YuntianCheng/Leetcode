//下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。
//
// 示例1:
//
//
// 输入：num = 2（或者0b10）
// 输出：[4, 1] 或者（[0b100, 0b1]）
//
//
// 示例2:
//
//
// 输入：num = 1
// 输出：[2, -1]
//
//
// 提示:
//
//
// num的范围在[1, 2147483647]之间；
// 如果找不到前一个或者后一个满足条件的正数，那么输出 -1。
//
//
// Related Topics 位运算 👍 56 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findClosedNumbers(num int) []int {
	if num == 2147483647 {
		return []int{-1, -1}
	}
	var a, b = 1, 2
	var count, sum1, sum2 int
	var res = []int{0, 0}
	for a <= num {
		sum2 = (sum2 << 1) + 1
		if a&num > 0 {
			sum1 += 1 << count
			count++
			if b&num == 0 && res[0] == 0 {
				tmp := sum2 & num
				res[0] = num - tmp + b + (sum1 >> 1)
			}
		}
		if b&num > 0 && a&num == 0 && res[1] == 0 {
			p, c := b, count
			res[1] = num - p - sum1
			for c >= 0 {
				res[1] = res[1] + (p >> 1)
				p >>= 1
				c--
			}
		}
		if res[0] != 0 && res[1] != 0 {
			break
		}
		a, b = b, b<<1
	}
	for i := range res {
		if res[i] == 0 {
			res[i] = -1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findClosedNumbers(124))
}
