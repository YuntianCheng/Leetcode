//字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没
//有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
//
// 示例1:
//
//
// 输入："aabcccccaaa"
// 输出："a2b1c5a3"
//
//
// 示例2:
//
//
// 输入："abbccd"
// 输出："abbccd"
// 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
//
//
// 提示：
//
//
// 字符串长度在[0, 50000]范围内。
//
//
// Related Topics 双指针 字符串 👍 147 👎 0

package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func compressString(S string) string {
	if S == "" {
		return S
	}
	s := strings.Builder{}
	var count = 1
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			count++
		} else {
			s.WriteByte(S[i-1])
			s.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	s.WriteByte(S[len(S)-1])
	s.WriteString(strconv.Itoa(count))
	if len(s.String()) < len(S) {
		return s.String()
	}
	return S
}

//leetcode submit region end(Prohibit modification and deletion)
