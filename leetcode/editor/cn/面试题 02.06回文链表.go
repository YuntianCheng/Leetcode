//编写一个函数，检查输入的链表是否是回文的。
//
//
//
// 示例 1：
//
// 输入： 1->2
//输出： false
//
//
// 示例 2：
//
// 输入： 1->2->2->1
//输出： true
//
//
//
//
// 进阶： 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 124 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var fast, slow = head, head
	var count int
	for fast.Next != nil {
		count++
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	var pre, cur = slow, slow.Next
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	for i := 0; i < count; i++ {
		if head.Val != fast.Val {
			return false
		}
		head = head.Next
		fast = fast.Next
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
