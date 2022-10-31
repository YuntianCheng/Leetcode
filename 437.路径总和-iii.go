/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */
package main

import (
	"fmt"
)

// @lc code=start
//   Definition for a binary tree node.

func countTreePath(root *TreeNode, sum int, target int) int {
	if root == nil {
		return 0
	}
	sum += root.Val
	if sum == target {
		return 1 + countTreePath(root.Left, sum, target) + countTreePath(root.Right, sum, target)
	} else {
		return countTreePath(root.Left, sum, target) + countTreePath(root.Right, sum, target)
	}
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := countTreePath(root, 0, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

// @lc code=end
func main() {
	root := &TreeNode{
		Val:  -2,
		Left: nil,
		Right: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(pathSum(root, -5))
}
