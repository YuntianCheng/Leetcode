//给定一个字符串 s 和一个字符串字典
// wordDict ，在字符串
// s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
//
// 注意：词典中的同一个单词可能在分段中被重复使用多次。
//
//
//
// 示例 1：
//
//
//输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
//输出:["cats and dog","cat sand dog"]
//
//
// 示例 2：
//
//
//输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine",
//"pineapple"]
//输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
//解释: 注意你可以重复使用字典中的单词。
//
//
// 示例 3：
//
//
//输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
//输出:[]
//
//
//
//
// 提示：
//
//
//
//
//
// 1 <= s.length <= 20
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 10
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中所有字符串都 不同
//
//
// Related Topics 字典树 记忆化搜索 哈希表 字符串 动态规划 回溯 👍 639 👎 0

package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
type Empty struct {
}

func dfs(s string, result *[]string, preFix string, wordDictMap map[string]Empty, maxLength int, dp *[]bool) {
	if !(*dp)[len(s)] {
		return
	}
	if s == "" {
		*result = append(*result, strings.TrimPrefix(preFix, " "))
		return
	}
	length := maxLength
	if length > len(s) {
		length = len(s)
	}
	originalLen := len(*result)
	for i := 0; i < length; i++ {
		if _, ok := wordDictMap[s[:i+1]]; ok {
			newPreFix := preFix + " " + s[:i+1]
			dfs(strings.TrimPrefix(s, s[:i+1]), result, newPreFix, wordDictMap, maxLength, dp)
		}
	}
	if originalLen == len(*result) {
		(*dp)[len(s)] = false
	}
}

func wordBreak(s string, wordDict []string) []string {
	wordDictMap := make(map[string]Empty, len(wordDict))
	maxLength := 0
	dp := make([]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = true
	}
	for _, word := range wordDict {
		wordDictMap[word] = Empty{}
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}
	result := make([]string, 0)
	dfs(s, &result, "", wordDictMap, maxLength, &dp)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
