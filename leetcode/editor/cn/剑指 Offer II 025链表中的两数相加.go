//给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 示例1：
//
//
//
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
//
//
// 示例2：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
//
//
// 示例3：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 100]
// 0 <= node.val <= 9
// 输入数据保证链表代表的数字无前导 0
//
//
//
//
// 进阶：如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。
//
//
//
//
// 注意：本题与主站 445 题相同：https://leetcode-cn.com/problems/add-two-numbers-ii/
//
// Related Topics 栈 链表 数学 👍 69 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cStack1, cStack2 := make([]*ListNode, 0), make([]*ListNode, 0)
	var top1, top2 = -1, -1
	var c1, c2 = l1, l2
	for c1 != nil {
		cStack1 = append(cStack1, c1)
		top1++
		c1 = c1.Next
	}
	for c2 != nil {
		cStack2 = append(cStack2, c2)
		top2++
		c2 = c2.Next
	}
	var sum, up int
	for top1 >= 0 || top2 >= 0 || up != 0 {
		sum = up
		if top1 >= 0 {
			sum += cStack1[top1].Val
		}
		if top2 >= 0 {
			sum += cStack2[top2].Val
		}
		up = sum / 10
		if top1 >= 0 || top2 >= 0 {
			if len(cStack1) >= len(cStack2) {
				cStack1[top1].Val = sum % 10
			} else {
				cStack2[top2].Val = sum % 10
			}
			top1--
			top2--
		} else {
			newNode := &ListNode{
				Val: sum,
			}
			if len(cStack1) >= len(cStack2) {
				newNode.Next = l1
				l1 = newNode
			} else {
				newNode.Next = l2
				l2 = newNode
			}
		}
	}
	if len(cStack1) >= len(cStack2) {
		return l1
	}
	return l2
}

//leetcode submit region end(Prohibit modification and deletion)
