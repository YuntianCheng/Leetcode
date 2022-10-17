//给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
//
//输入: a = "11", b = "10"
//输出: "101"
//
// 示例 2:
//
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//
// 提示：
//
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
//
//
//
//
//
// 注意：本题与主站 67 题相同：https://leetcode-cn.com/problems/add-binary/
//
// Related Topics 位运算 数学 字符串 模拟 👍 49 👎 0

package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	var o, t = a, b
	if len(b) > len(a) {
		o, t = b, a
	}
	var up uint8 = 0
	result := make([]uint8, len(o)+1)
	var i, j, k = len(o) - 1, len(t) - 1, len(result) - 1
	for j >= 0 {
		sum := o[i] + t[j] + up - '0'
		if sum > '1' {
			up = 1
			sum = sum - '2' + '0'
		} else {
			up = 0
		}
		result[k] = sum
		k--
		i--
		j--
	}
	for i >= 0 {
		sum := o[i] + up
		if sum > '1' {
			up = 1
			sum = sum - '2' + '0'
		} else {
			up = 0
		}
		result[k] = sum
		i--
		k--
	}
	if up == 1 {
		result[k] = '1'
		return string(result)
	}
	return string(result[1:])
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(addBinary("11", "10"))
}
