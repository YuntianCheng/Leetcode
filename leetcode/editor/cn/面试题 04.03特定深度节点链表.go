//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
//
//
// 示例：
//
// 输入：[1,2,3,4,5,null,7,8]
//
//        1
//       /  \
//      2    3
//     / \    \
//    4   5    7
//   /
//  8
//
//输出：[[1],[2,3],[4,5,7],[8]]
//
//
// Related Topics 树 广度优先搜索 链表 二叉树 👍 80 👎 0

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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func dfs0403(nodes []*TreeNode, result *[]*ListNode) {
	length := len(nodes)
	fake := &ListNode{}
	cur := fake
	for i := 0; i < length; i++ {
		cur.Next = &ListNode{
			Val: nodes[i].Val,
		}
		cur = cur.Next
		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
		}
	}
	if fake.Next != nil {
		*result = append(*result, fake.Next)
	}
	if len(nodes) > length {
		dfs0403(nodes[length:], result)
	}
}

func listOfDepth(tree *TreeNode) []*ListNode {
	nodes := make([]*TreeNode, 1)
	nodes[0] = tree
	result := make([]*ListNode, 0)
	dfs0403(nodes, &result)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
