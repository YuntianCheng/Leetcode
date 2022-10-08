//字符串有三种编辑操作:插入一个英文字符、删除一个英文字符或者替换一个英文字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
// 示例 1:
//
//
//输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//
// 示例 2:
//
//
//输入:
//first = "pales"
//second = "pal"
//输出: False
//
//
// Related Topics 双指针 字符串 👍 227 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	var i, j, lF, lS, change = 0, 0, len(first), len(second), 0
	for i < len(first) && j < len(second) {
		if first[i] == second[j] {
			i++
			j++
			lF--
			lS--
		} else {
			if change > 0 {
				return false
			}
			change++
			if lS < lF {
				i++
				lF--
			} else if lS == lF {
				i++
				j++
				lF--
				lS--
			} else {
				j++
				lS--
			}
		}
	}

	return (lF == 0 && lS == 0) || (lF <= 1 && lS <= 1 && change == 0)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(oneEditAway("", "a"))
}
