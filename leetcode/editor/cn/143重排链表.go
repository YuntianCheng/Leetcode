//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
//L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [1,2,3,4]
//输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5]
//输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
//
// Related Topics 栈 递归 链表 双指针 👍 1032 👎 0

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

func revert(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	current := pre.Next
	pre.Next = nil
	for current != nil {
		tmp := current.Next
		current.Next = pre
		pre = current
		current = tmp
	}
	return pre
}
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	oneStep := head
	twoStep := head.Next
	for twoStep != nil && twoStep.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
	}
	insert := revert(oneStep.Next)
	oneStep.Next = nil
	newHead := head
	for insert != nil {
		tmp1 := newHead.Next
		tmp2 := insert
		newHead.Next = tmp2
		insert = insert.Next
		tmp2.Next = tmp1
		newHead = tmp1
	}
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
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
