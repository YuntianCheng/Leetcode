//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 1798 👎 0

package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode, nextStart *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	current := head.Next
	//pre.Next = nextStart
	for current != nil {
		tmp := current.Next
		current.Next = pre
		pre = current
		current = tmp
	}
	head.Next = nextStart
	return head
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	count := 0
	start := head
	current := head
	var preEnd *ListNode
	var hasHead bool
	for current != nil {
		count++
		if count == k {
			tmp := current.Next
			current.Next = nil
			end := reverseList(start, tmp)
			if hasHead {
				preEnd.Next = current
			}
			if !hasHead {
				head = current
				hasHead = true
			}
			preEnd = end
			start = tmp
			current = tmp
			count = 0
		} else {
			current = current.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	newHead := reverseKGroup(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
