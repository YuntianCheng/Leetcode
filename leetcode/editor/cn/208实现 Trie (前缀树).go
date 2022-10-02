//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼
//写检查。
//
// 请你实现 Trie 类：
//
//
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回
//false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否
//则，返回 false 。
//
//
//
//
// 示例：
//
//
//输入
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
//
//解释
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 True
//trie.search("app");     // 返回 False
//trie.startsWith("app"); // 返回 True
//trie.insert("app");
//trie.search("app");     // 返回 True
//
//
//
//
// 提示：
//
//
// 1 <= word.length, prefix.length <= 2000
// word 和 prefix 仅由小写英文字母组成
// insert、search 和 startsWith 调用次数 总计 不超过 3 * 10⁴ 次
//
//
// Related Topics 设计 字典树 哈希表 字符串 👍 1310 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	trees map[uint8]*Tree
}

type Tree struct {
	isWord   bool
	children []*Tree
	char     uint8
}

func Constructor() Trie {
	trie := Trie{
		trees: make(map[uint8]*Tree, 0),
	}
	return trie
}

func (this *Trie) Insert(word string) {
	var node *Tree
	for i := 0; i < len(word); i++ {
		if i == 0 {
			if v, ok := this.trees[word[i]]; ok {
				node = v
			} else {
				node = &Tree{
					char:     word[i],
					children: make([]*Tree, 0),
				}
				this.trees[word[0]] = node
			}
		} else {
			var find bool
			for j := 0; j < len(node.children); j++ {
				if node.children[j].char == word[i] {
					find = true
					node = node.children[j]
					break
				}
			}
			if !find {
				tmp := &Tree{
					char:     word[i],
					children: make([]*Tree, 0),
				}
				node.children = append(node.children, tmp)
				node = tmp
			}
		}
		if i == len(word)-1 {
			node.isWord = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	if v, ok := this.trees[word[0]]; ok {
		tmp := v
		for i := 1; i < len(word); i++ {
			var find bool
			for j := 0; j < len(tmp.children); j++ {
				if tmp.children[j].char == word[i] {
					tmp = tmp.children[j]
					find = true
					break
				}
			}
			if !find {
				return false
			}
		}
		if tmp.isWord {
			return true
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if v, ok := this.trees[prefix[0]]; ok {
		tmp := v
		for i := 1; i < len(prefix); i++ {
			var find bool
			for j := 0; j < len(tmp.children); j++ {
				if tmp.children[j].char == prefix[i] {
					tmp = tmp.children[j]
					find = true
					break
				}
			}
			if !find {
				return false
			}
		}
		return true
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	trie := Constructor()
	trie.Insert("a")
	fmt.Println(trie.Search("a"))
	fmt.Println(trie.StartsWith("a"))
}
