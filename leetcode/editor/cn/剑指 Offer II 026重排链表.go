//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln-1 → Ln 请将其重新排列后变为：
//
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1:
//
//
//
//
//输入: head = [1,2,3,4]
//输出: [1,4,2,3]
//
// 示例 2:
//
//
//
//
//输入: head = [1,2,3,4,5]
//输出: [1,5,2,4,3]
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
//
//
//
// 注意：本题与主站 143 题相同：https://leetcode-cn.com/problems/reorder-list/
//
// Related Topics 栈 递归 链表 双指针 👍 78 👎 0

package main

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
func reorderList(head *ListNode) {
	var slow, fast = head, head
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	var pre = slow.Next
	slow.Next = nil
	var cur *ListNode
	if pre != nil {
		cur = pre.Next
		pre.Next = nil
		for cur != nil {
			pre, cur, cur.Next = cur, cur.Next, pre
		}
		cur = head
		for fast != nil {
			cur.Next, fast.Next, cur, fast = fast, cur.Next, cur.Next, fast.Next
		}
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
					Val: 4,
				},
			},
		},
	}
	reorderList(head)
}
