//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,2,1]
//输出：true
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 1526 👎 0

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
func isPalindrome234(head *ListNode) bool {
	mid, end := head, head
	var count = 1
	for end.Next != nil {
		end = end.Next
		count++
		if end.Next != nil {
			count++
			end = end.Next
		} else {
			break
		}
		mid = mid.Next
	}
	current := mid
	next := current.Next
	current.Next = mid
	for current != end {
		tmp := next.Next
		next.Next = current
		current = next
		next = tmp
	}
	for i := 0; i < count/2; i++ {
		if head.Val != end.Val {
			return false
		}
		head = head.Next
		end = end.Next
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))
}
