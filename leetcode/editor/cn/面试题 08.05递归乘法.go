//递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。
//
// 示例1:
//
//
// 输入：A = 1, B = 10
// 输出：10
//
//
// 示例2:
//
//
// 输入：A = 3, B = 4
// 输出：12
//
//
// 提示:
//
//
// 保证乘法范围不会溢出
//
//
// Related Topics 位运算 递归 数学 👍 76 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}
	if A == 1 {
		return B
	}
	if B == 1 {
		return A
	}
	if A > B {
		return A + multiply(A, B-1)
	} else {
		return B + multiply(A-1, B)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
