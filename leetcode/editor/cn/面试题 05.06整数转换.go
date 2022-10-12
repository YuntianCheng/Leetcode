//整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。
//
// 示例1:
//
//
// 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
// 输出：2
//
//
// 示例2:
//
//
// 输入：A = 1，B = 2
// 输出：2
//
//
// 提示:
//
//
// A，B范围在[-2147483648, 2147483647]之间
//
//
// Related Topics 位运算 👍 46 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func convertInteger(A int, B int) int {
	C := uint32(A ^ B)
	var count int
	for C != 0 {
		if C&1 > 0 {
			count++
		}
		C >>= 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(convertInteger(826966453, -729934991))
}
