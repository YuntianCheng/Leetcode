//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2195 👎 0

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
func addElm(heap *[]*ListNode, elm *ListNode) {
	*heap = append(*heap, elm)
	var i = len(*heap) - 1
	for i != 0 {
		if (*heap)[i].Val < (*heap)[(i-1)/2].Val {
			(*heap)[(i-1)/2], (*heap)[i] = (*heap)[i], (*heap)[(i-1)/2]
			i = (i - 1) / 2
		} else {
			break
		}
	}
}

func getMin(heap *[]*ListNode) *ListNode {
	if len(*heap) > 0 {
		return (*heap)[0]
	}
	return nil
}

func pop(heap *[]*ListNode) {
	if len(*heap) == 0 {
		return
	}
	(*heap)[0], (*heap)[len(*heap)-1] = (*heap)[len(*heap)-1], (*heap)[0]
	*heap = (*heap)[:len(*heap)-1]
	var i = 0
	for 2*i+1 < len(*heap) {
		var k = 2*i + 1
		if 2*i+2 < len(*heap) && (*heap)[2*i+2].Val < (*heap)[k].Val {
			k = 2*i + 2
		}
		if (*heap)[i].Val > (*heap)[k].Val {
			(*heap)[k], (*heap)[i] = (*heap)[i], (*heap)[k]
			i = k
		} else {
			break
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var heap []*ListNode
	h := lists[0]
	for h != nil {
		heap = append(heap, h)
		h = h.Next
	}
	for i := 1; i < len(lists); i++ {
		c := lists[i]
		for c != nil {
			addElm(&heap, c)
			c = c.Next
		}
	}
	if len(heap) == 0 {
		return nil
	}
	head, cur := getMin(&heap), getMin(&heap)
	pop(&heap)
	for len(heap) > 0 {
		cur.Next = getMin(&heap)
		pop(&heap)
		cur = cur.Next
	}
	cur.Next = nil
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	//head1 := &ListNode{
	//	1,
	//	&ListNode{
	//		4,
	//		&ListNode{
	//			5, nil,
	//		},
	//	},
	//}
	//head2 := &ListNode{
	//	1,
	//	&ListNode{
	//		3,
	//		&ListNode{
	//			4, nil,
	//		},
	//	},
	//}
	//head3 := &ListNode{
	//	2,
	//	&ListNode{
	//		6,
	//		nil,
	//	},
	//}
	fmt.Println(mergeKLists([]*ListNode{nil}))
}
