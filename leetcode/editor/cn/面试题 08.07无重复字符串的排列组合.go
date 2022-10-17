//无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
//
// 示例1:
//
//
// 输入：S = "qwe"
// 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
//
//
// 示例2:
//
//
// 输入：S = "ab"
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
// Related Topics 字符串 回溯 👍 81 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func permutation1(S string) []string {
	var result []string
	var dfs87 func(string, string, int)
	dfs87 = func(S string, val string, index int) {
		val += string(S[index])
		K := "" + S[:index] + S[index+1:]
		if K == "" {
			result = append(result, val)
			return
		}
		for i := range K {
			dfs87(K, val, i)
		}
	}
	for i := 0; i < len(S); i++ {
		dfs87(S, "", i)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(permutation("qwe"))
}
