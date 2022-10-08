//设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
//
// 如果指定节点没有对应的“下一个”节点，则返回null。
//
// 示例 1:
//
// 输入: root = [2,1,3], p = 1
//
//  2
// / \
//1   3
//
//输出: 2
//
// 示例 2:
//
// 输入: root = [5,3,6,2,4,null,null,1], p = 6
//
//      5
//     / \
//    3   6
//   / \
//  2   4
// /
//1
//
//输出: null
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 209 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var result *TreeNode
	stack := make([]*TreeNode, 0)
	cur := root
	top := -1
	var is bool
	for top >= 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			top++
			cur = cur.Left
		}
		if is {
			result = stack[top]
			break
		}
		if stack[top] == p {
			is = true
		}
		cur = stack[top].Right
		stack = stack[:top]
		top--
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
