//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
//
//
//
// 示例 1:
//
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//
// 示例 2:
//
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 200
// num1 和 num2 只能由数字组成。
// num1 和 num2 都不包含任何前导零，除了数字0本身。
//
//
// Related Topics 数学 字符串 模拟 👍 1056 👎 0

package main

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func add(num1 string, num2 string) string {
	var result []string
	if len(num1) > len(num2) {
		result = make([]string, len(num1)+1)
	} else {
		result = make([]string, len(num2)+1)
	}
	k := len(num1) - 1
	l := len(num2) - 1
	up := 0
	for i := len(result) - 1; i >= 0; i-- {
		var a, b int
		if k >= 0 {
			a, _ = strconv.Atoi(string(num1[k]))
		} else {
			a = 0
		}
		if l >= 0 {
			b, _ = strconv.Atoi(string(num2[l]))
		} else {
			b = 0
		}
		h := strconv.Itoa(a + b + up)
		if len(h) > 1 {
			result[i] = h[1:]
			up = 1
		} else {
			if i != 0 {
				result[i] = h
				up = 0
			} else if up == 1 {
				result[i] = "1"
			}
		}
		l--
		k--
	}
	return strings.Join(result, "")
}
func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	if len(num1) < 3 && len(num2) < 3 {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		return strconv.Itoa(a * b)
	}
	len1 := len(num1) / 2
	len2 := len(num2) / 2
	num1F := num1[:len1]
	if num1F == "" {
		num1F = "0"
	}
	num1S := strings.TrimLeft(num1[len1:], "0")
	if num1S == "" {
		num1S = "0"
	}
	num2F := num2[:len2]
	if num2F == "" {
		num2F = "0"
	}
	num2S := strings.TrimLeft(num2[len2:], "0")
	if num2S == "" {
		num2S = "0"
	}
	tmp := ""
	for i := 0; i < len(num1)-len1; i++ {
		tmp += "0"
	}
	num1FM := tmp
	tmp = ""
	for i := 0; i < len(num2)-len2; i++ {
		tmp += "0"
	}
	num2FM := tmp
	x1 := multiply(num1F, num2F)
	if x1 != "0" {
		x1 += num1FM + num2FM
	}
	x2 := multiply(num1S, num2F)
	if x2 != "0" {
		x2 += num2FM
	}
	x3 := multiply(num1F, num2S)
	if x3 != "0" {
		x3 += num1FM
	}
	return add(add(x1, x2), add(x3, multiply(num1S, num2S)))
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(multiply("992", "1"))
}
