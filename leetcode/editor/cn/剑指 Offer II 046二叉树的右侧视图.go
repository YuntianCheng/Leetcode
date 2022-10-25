//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//
//
// 示例 1:
//
//
//
//
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]
//
//
// 示例 2:
//
//
//输入: [1,null,3]
//输出: [1,3]
//
//
// 示例 3:
//
//
//输入: []
//输出: []
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,100]
//
// -100 <= Node.val <= 100
//
//
//
//
//
// 注意：本题与主站 199 题相同：https://leetcode-cn.com/problems/binary-tree-right-side-
//view/
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 37 👎 0

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var queue = []*TreeNode{root}
	var start, end, length = 0, 0, 1
	var result = make([]int, 0)
	for start < length {
		node := queue[start]
		result = append(result, node.Val)
		var newEnd = end
		for start <= end {
			if queue[start].Right != nil {
				queue = append(queue, queue[start].Right)
				newEnd++
				length++
			}
			if queue[start].Left != nil {
				queue = append(queue, queue[start].Left)
				newEnd++
				length++
			}
			start++
		}
		end = newEnd
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
