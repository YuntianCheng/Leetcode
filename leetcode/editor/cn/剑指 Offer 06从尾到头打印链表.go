//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
// 示例 1：
//
// 输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//
// 限制：
//
// 0 <= 链表长度 <= 10000
//
// Related Topics 栈 递归 链表 双指针 👍 347 👎 0

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	c := head
	var p *ListNode
	for c.Next != nil {
		p, c, c.Next = c, c.Next, p
	}
	c.Next = p
	var r []int
	for c != nil {
		r = append(r, c.Val)
		c = c.Next
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
