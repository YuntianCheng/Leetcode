//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 2673 👎 0

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	point1 := list1
	point2 := list2
	var pre *ListNode
	for point2 != nil && point1 != nil {
		if point1.Val <= point2.Val {
			pre = point1
			point1 = point1.Next
		} else {
			pre.Next = point2
			pre = pre.Next
			tmp := point2.Next
			point2.Next = point1
			point2 = tmp
		}
	}
	if point1 == nil {
		if pre != nil {
			pre.Next = point2
		} else {
			return list2
		}
	}
	return list1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(list1, list2))
}
