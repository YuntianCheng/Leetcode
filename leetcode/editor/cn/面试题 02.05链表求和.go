//给定两个用链表表示的整数，每个节点包含一个数位。
//
// 这些数位是反向存放的，也就是个位排在链表首部。
//
// 编写函数对这两个整数求和，并用链表形式返回结果。
//
//
//
// 示例：
//
// 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//
//
// 进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
// 示例：
//
// 输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//
//
// Related Topics 递归 链表 数学 👍 149 👎 0

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1, c2 := l1, l2
	var i int
	var pre *ListNode
	for c1 != nil && c2 != nil {
		sum := c1.Val + c2.Val + i
		i, c2.Val = sum/10, sum%10
		c1, c2, pre = c1.Next, c2.Next, c2
	}
	for c1 != nil {
		sum := c1.Val + i
		i = sum / 10
		v := sum % 10
		pre.Next = &ListNode{
			v, nil,
		}
		pre, c1 = pre.Next, c1.Next
	}
	for c2 != nil && i > 0 {
		sum := c2.Val + i
		i, c2.Val = sum/10, sum%10
		c2, pre = c2.Next, c2
	}
	if i > 0 {
		pre.Next = &ListNode{
			1, nil,
		}
	}
	return l2
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	l1 := &ListNode{
		1, &ListNode{
			8, nil,
		},
	}
	l2 := &ListNode{
		0, nil,
	}
	addTwoNumbers(l1, l2)
}
