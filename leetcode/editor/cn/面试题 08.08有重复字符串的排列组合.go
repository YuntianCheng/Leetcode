//有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
//
// 示例1:
//
//  输入：S = "qqe"
// 输出：["eqq","qeq","qqe"]
//
//
// 示例2:
//
//  输入：S = "ab"
// 输出：["ab", "ba"]
//
//
// 提示:
//
//
// 字符都是英文字母。
// 字符串长度在[1, 9]之间。
//
//
// Related Topics 字符串 回溯 👍 78 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func permutation(S string) []string {
	var result []string
	var dfs88 func([]byte, int)
	dfs88 = func(S []byte, index int) {
		if index == len(S)-1 {
			result = append(result, string(S))
		}
		var set uint
		for j := index; j < len(S); j++ {
			if set&(1<<(S[j]-'A')) == 0 {
				set |= 1 << (S[j] - 'A')
				S[index], S[j] = S[j], S[index]
				dfs88(S, index+1)
				S[index], S[j] = S[j], S[index]
			}
		}
	}
	dfs88([]byte(S), 0)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(permutation("qqe"))
}
