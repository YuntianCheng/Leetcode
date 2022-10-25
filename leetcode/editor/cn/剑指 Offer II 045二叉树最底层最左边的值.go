//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [2,1,3]
//输出: 1
//
//
// 示例 2:
//
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
//
// 注意：本题与主站 513 题相同： https://leetcode-cn.com/problems/find-bottom-left-tree-
//value/
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 32 👎 0

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
func findBottomLeftValue(root *TreeNode) int {
	var ddd func(*TreeNode, int) (int, int)
	ddd = func(root *TreeNode, deep int) (val int, level int) {
		if root.Left == nil && root.Right == nil {
			return root.Val, deep + 1
		}
		var lv, rv, ll, rl int
		if root.Left != nil {
			lv, ll = ddd(root.Left, deep+1)
		}
		if root.Right != nil {
			rv, rl = ddd(root.Right, deep+1)
		}
		if rl > ll {
			return rv, rl
		}
		return lv, ll
	}
	v, _ := ddd(root, 0)
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
