/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
 */
package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	h := make(map[int]int, 0)
	result := make([]int, 0)
	getTreeNodeSum(root, h)
	max := 0
	for _, val := range h {
		if val > max {
			max = val
		}
	}
	for key := range h {
		if h[key] == max {
			result = append(result, key)
		}
	}
	return result
}
func getTreeNodeSum(root *TreeNode, nums map[int]int) (sum int) {
	me := root.Val
	if root.Left != nil {
		me = me + getTreeNodeSum(root.Left, nums)
	}
	if root.Right != nil {
		me = me + getTreeNodeSum(root.Right, nums)
	}
	if _, ok := nums[me]; ok {
		nums[me]++
	} else {
		nums[me] = 1
	}
	return me
}

// @lc code=end
func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(findFrequentTreeSum(root))
}
