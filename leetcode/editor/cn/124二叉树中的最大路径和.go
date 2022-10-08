//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不
//一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//
//
// 提示：
//
//
// 树中节点数目范围是 [1, 3 * 10⁴]
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 1740 👎 0

package main

import "fmt"

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

func dfs124(root *TreeNode, m map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	a := dfs124(root.Left, m)
	b := dfs124(root.Right, m)
	result, path := root.Val, root.Val
	if a > 0 {
		path += a
		if a > b {
			result += a
		}
	}
	if b > 0 {
		path += b
		if b > a {
			result += b
		}
	}
	if _, ok := m[root]; !ok {
		m[root] = path
	}
	return result
}

func getMaxPathValue(root *TreeNode, max *int, m map[*TreeNode]int) {
	if root == nil {
		return
	}
	var path int
	if v, ok := m[root]; ok {
		path = v
	} else {
		a := dfs124(root.Left, m)
		b := dfs124(root.Right, m)
		path += root.Val
		if a > 0 {
			path += a
		}
		if b > 0 {
			path += b
		}
	}
	if path > *max {
		*max = path
	}
	getMaxPathValue(root.Left, max, m)
	getMaxPathValue(root.Right, max, m)
}

func maxPathSum(root *TreeNode) int {
	m := make(map[*TreeNode]int, 0)
	var max = ^int(^uint(0) >> 1)
	getMaxPathValue(root, &max, m)
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: -1,
		},
	}
	fmt.Println(maxPathSum(root))
}
