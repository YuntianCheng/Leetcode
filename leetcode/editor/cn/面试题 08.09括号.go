//括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
//
// 说明：解集不能包含重复的子集。
//
// 例如，给出 n = 3，生成结果为：
//
//
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
//
//
// Related Topics 字符串 动态规划 回溯 👍 119 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	var result []string
	var dfssss func(int, int, string)
	dfssss = func(left, right int, block string) {
		if left == 0 && right == 0 {
			result = append(result, block)
			return
		}
		if left > 0 {
			dfssss(left-1, right, block+"(")
		}
		if right > left {
			dfssss(left, right-1, block+")")
		}
	}
	dfssss(n, n, "")
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
