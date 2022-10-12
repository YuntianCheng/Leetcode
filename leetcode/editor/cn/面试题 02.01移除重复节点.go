//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
//
// Related Topics 哈希表 链表 双指针 👍 166 👎 0

package main

import "fmt"

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
const l = 20000/64 + 1

func removeDuplicateNodes(head *ListNode) *ListNode {
	var a [l]uint64
	front := &ListNode{
		Next: head,
	}
	pre := front
	for head != nil {
		var t uint64 = 1 << (head.Val % 64)
		if t&a[head.Val/64] > 0 {
			pre.Next = head.Next
		} else {
			pre = pre.Next
		}
		a[head.Val/64] = a[head.Val/64] | t
		head = head.Next
	}
	return front.Next
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
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Println(removeDuplicateNodes(head))
}
