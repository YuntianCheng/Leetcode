//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
// Related Topics 链表 双指针 分治 排序 归并排序 👍 1793 👎 0

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

func mergeSort(start *ListNode, end *ListNode) (*ListNode, *ListNode) {
	if start != end {
		quick, slow := start, start
		for quick != end {
			quick = quick.Next
			if quick != end {
				quick = quick.Next
			} else {
				break
			}
			slow = slow.Next
		}
		start, slow = mergeSort(start, slow)
		var stop1 *ListNode
		stop1, end = mergeSort(slow.Next, end)
		slow.Next = stop1
		start1 := start
		start2 := stop1
		stop2 := end.Next
		var head *ListNode
		for start1 != stop1 && start2 != stop2 {
			var current *ListNode
			if start1.Val > start2.Val {
				current = start2
				start2 = start2.Next
			} else {
				current = start1
				start1 = start1.Next
			}
			if head == nil {
				head = current
				start = head
			} else {
				head.Next = current
				head = head.Next
			}
		}
		for start1 != stop1 {
			head.Next = start1
			start1 = start1.Next
			head = head.Next
		}
		for start2 != stop2 {
			head.Next = start2
			start2 = start2.Next
			head = head.Next
		}
		end = head
		head.Next = stop2
		//fmt.Println(head)
	}
	return start, end
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	end := head
	for end.Next != nil {
		end = end.Next
	}
	start, _ := mergeSort(head, end)
	return start
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	head := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(sortList(head))
}
