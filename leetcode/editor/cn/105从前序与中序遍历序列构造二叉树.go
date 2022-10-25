//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
//返回其根节点。
//
//
//
// 示例 1:
//
//
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
//
//
//
//
// 提示:
//
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1760 👎 0

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
	var dfsTree func(int, int, int, *TreeNode)
	dfsTree = func(rootIndex, start, end int, root *TreeNode) {
		if start > end {
			return
		}
		root.Val = preorder[rootIndex]
		if start == end {
			return
		}
		var i = start
		for preorder[rootIndex] != inorder[i] && i < end {
			i++
		}
		if i != start {
			root.Left = &TreeNode{}
			dfsTree(rootIndex+1, start, i-1, root.Left)
		}
		if i != end {
			root.Right = &TreeNode{}
			dfsTree(rootIndex+1+i-start, i+1, end, root.Right)
		}

	}
	root := &TreeNode{}
	dfsTree(0, 0, len(inorder)-1, root)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
}
