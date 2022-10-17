//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
// 返回被除数 dividend 除以除数 divisor 得到的商。
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
// 示例 1:
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//
// 示例 2:
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2
//
//
//
// 提示：
//
//
// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2³¹, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
//
// Related Topics 位运算 数学 👍 991 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func divide2(dividend int, divisor int) int {
	dvd, dvs := int32(dividend), int32(divisor)
	if dvd == ^int32(^uint32(0)>>1) && dvs == -1 {
		return int(int32(^uint32(0) >> 1))
	}
	if dvs == ^int32(^uint32(0)>>1) {
		if dvd > ^int32(^uint32(0)>>1) {
			return 0
		}
		return 1
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}
	var k = 1
	result := 0
	if (dvd > 0 && dvs < 0) || (dvd < 0 && dvs > 0) {
		k = -1
	}

	if dvs < 0 {
		dvs = -dvs
	}
	if dvd < 0 {
		if dvd == ^int32(^uint32(0)>>1) {
			dvd += dvs
			result++
		}
		dvd = -dvd
	}
	for dvd >= dvs {
		tmp := dvs
		for i := 1; dvd >= tmp && tmp > 0; tmp, i = tmp<<1, i<<1 {
			dvd -= tmp
			result += i
		}
	}
	return result * k
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(divide(-2147483648, 2))
}
