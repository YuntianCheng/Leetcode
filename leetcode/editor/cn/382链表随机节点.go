//给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。
//
// 实现 Solution 类：
//
//
// Solution(ListNode head) 使用整数数组初始化对象。
// int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。
//
//
//
//
// 示例：
//
//
//输入
//["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
//[[[1, 2, 3]], [], [], [], [], []]
//输出
//[null, 1, 3, 2, 2, 3]
//
//
//解释
//Solution solution = new Solution([1, 2, 3]);
//solution.getRandom(); // 返回 1
//solution.getRandom(); // 返回 3
//solution.getRandom(); // 返回 2
//solution.getRandom(); // 返回 2
//solution.getRandom(); // 返回 3
//// getRandom() 方法应随机返回 1、2、3中的一个，每个元素被返回的概率相等。
//
//
//
// 提示：
//
//
// 链表中的节点数在范围 [1, 10⁴] 内
// -10⁴ <= Node.val <= 10⁴
// 至多调用 getRandom 方法 10⁴ 次
//
//
//
//
// 进阶：
//
//
// 如果链表非常大且长度未知，该怎么处理？
// 你能否在不使用额外空间的情况下解决此问题？
//
//
// Related Topics 水塘抽样 链表 数学 随机化 👍 314 👎 0

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
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
type Solution struct {
	Head *ListNode
	//Position []*ListNode
	//Num      int
}

//func Constructor(head *ListNode) Solution {
//	//var curosr, i = head, 0
//	//var s = Solution{
//	//	Head:     head,
//	//	Position: []*ListNode{head},
//	//}
//	//for curosr != nil {
//	//	i++
//	//	curosr = curosr.Next
//	//	if i%10 == 0 && curosr != nil {
//	//		s.Position = append(s.Position, curosr)
//	//	}
//	//}
//	//s.Num = i
//	return Solution{head}
//}

func (this *Solution) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	var result int
	for i, curosr := 1, this.Head; curosr != nil; i, curosr = i+1, curosr.Next {
		if rand.Intn(i) == 0 {
			result = curosr.Val
		}
	}
	return result
	//index := rand.Intn(this.Num + 1)
	//var j = 0
	//for index > 10 {
	//	j++
	//	index -= 10
	//}
	//cursor := this.Position[j]
	//for i := 1; i < index; i++ {
	//	cursor = cursor.Next
	//}
	//return cursor.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	head := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, nil,
			},
		},
	}
	s := Constructor(head)
	fmt.Println(s.GetRandom())
}
