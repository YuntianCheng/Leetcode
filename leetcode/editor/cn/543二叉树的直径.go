//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
//
//
// 示例 : 给定二叉树
//
//           1
//         / \
//        2   3
//       / \
//      4   5
//
//
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。
//
// Related Topics 树 深度优先搜索 二叉树 👍 1167 👎 0

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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getHeight(root *TreeNode, m map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if height, ok := m[root]; ok {
		return height
	}
	height := Max(getHeight(root.Left, m), getHeight(root.Right, m)) + 1
	m[root] = height
	return height
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := make(map[*TreeNode]int, 0)
	max := 0
	q := make([]*TreeNode, 0)
	q = append(q, root)
	start := 0
	end := 1
	for start < end {
		node := q[start]
		start++
		path := getHeight(node.Left, m) + getHeight(node.Right, m)
		if path > max {
			max = path
		}
		if node.Left != nil {
			q = append(q, node.Left)
			end++
		}
		if node.Right != nil {
			q = append(q, node.Right)
			end++
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}
