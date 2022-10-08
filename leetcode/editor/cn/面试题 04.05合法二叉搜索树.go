//实现一个函数，检查一棵二叉树是否为二叉搜索树。
//示例 1:
// 输入:     2    / \   1   3 输出: true
//示例 2:
// 输入:     5    / \   1   4      / \     3   6 输出: false 解释: 输入为: [5,1,4,null,
//null,3,6]。   根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 89 👎 0

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
func dfs0405(root *TreeNode, low, up int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= up {
		return false
	}
	return dfs0405(root.Left, low, root.Val) && dfs0405(root.Right, root.Val, up)
}

func isValidBST(root *TreeNode) bool {
	return dfs0405(root, ^int(^uint(0)>>1), int(^uint(0)>>1))
}

//leetcode submit region end(Prohibit modification and deletion)
