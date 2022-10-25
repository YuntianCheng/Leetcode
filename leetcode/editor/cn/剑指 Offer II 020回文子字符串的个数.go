//给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
//
//
//
// 注意：本题与主站 647 题相同：https://leetcode-cn.com/problems/palindromic-substrings/
//
// Related Topics 字符串 动态规划 👍 71 👎 0

package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countSubstrings(s string) int {
	build := strings.Builder{}
	build.WriteString("$#")
	for i := range s {
		build.Write([]byte{s[i], '#'})
	}
	build.WriteByte('!')
	newS := build.String()
	var f = make([]int, len(newS)-1)
	var count, rmax, imax = 0, 0, 0
	for i := 1; i < len(f); i++ {
		if i <= rmax {
			f[i] = min(rmax-i+1, f[2*imax-i])
		} else {
			f[i] = 1
		}
		for newS[i+f[i]] == newS[i-f[i]] {
			f[i]++
		}
		if i+f[i]-1 > rmax {
			rmax = i + f[i] - 1
			imax = i
		}
		count += f[i] / 2
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(countSubstrings("abc"))
}
