//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
//
// 示例 3：
//
//
//输入：root = []
//输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 1145 👎 0

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
//var (
//	m = make(map[*TreeNode]int, 0)
//)

func getHeight1(root *TreeNode, m map[*TreeNode]int) int {
	var height int
	if root != nil {
		if v, ok := m[root]; ok {
			return v
		}
		l := getHeight(root.Left, m)
		r := getHeight(root.Right, m)
		if l >= r {
			height = l + 1
		} else {
			height = r + 1
		}
	}
	m[root] = height
	return height
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := getHeight(root.Left, m)
	right := getHeight(root.Right, m)
	isLeft, isRight := isBalanced(root.Left), isBalanced(root.Right)
	return left-right <= 1 && left-right >= -1 && isLeft && isRight
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(isBalanced(root))
}
