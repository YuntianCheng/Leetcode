//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你不需要 保留 每个分区中各节点的初始相对位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200
//
//
// Related Topics 链表 双指针 👍 116 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	fake1 := &ListNode{}
	fake1.Next = head
	fake2 := &ListNode{}
	large, small := fake1, fake2
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = fake1.Next
	return fake2.Next
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	head := &ListNode{
		1, &ListNode{
			4, &ListNode{
				3, &ListNode{
					2, &ListNode{
						5, &ListNode{
							2, nil,
						},
					},
				},
			},
		},
	}
	partition(head, 3)
}
