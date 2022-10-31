//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
// 示例 1:
//
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000
//
//
//
// 注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-
//preorder-and-inorder-traversal/
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 930 👎 0

package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	var b func(int, int, int) *TreeNode
	b = func(rootIndex, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		root := &TreeNode{
			Val: preorder[rootIndex],
		}
		if start == end {
			return root
		}
		var i = start
		for ; i <= end; i++ {
			if inorder[i] == preorder[rootIndex] {
				break
			}
		}
		if i > start {
			root.Left = b(rootIndex+1, start, i-1)
		}
		if i < end {
			root.Right = b(rootIndex+i-start+1, i+1, end)
		}
		return root
	}
	return b(0, 0, len(inorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
