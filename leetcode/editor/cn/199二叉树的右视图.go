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
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 760 👎 0

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

//func result(root *TreeNode, r *[]int) {
//	*r = append(*r, root.Val)
//	if root.Right != nil {
//
//	}
//}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	start := 0
	end := 1
	count := start + 1
	for start < end {
		var i = start
		tmpCount := 0
		for ; i < count; i++ {
			tmp := q[i]
			if i == start {
				result = append(result, tmp.Val)
			}
			if tmp.Right != nil {
				q = append(q, tmp.Right)
				tmpCount++
				end++
			}
			if tmp.Left != nil {
				q = append(q, tmp.Left)
				tmpCount++
				end++
			}
		}
		start = i
		count = start + tmpCount
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		1, &TreeNode{
			2, nil, nil,
		}, nil,
	}
	fmt.Println(rightSideView(root))
}
