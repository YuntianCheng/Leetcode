//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：
//push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
//
// 示例1:
//
//  输入：
//["SortedStack", "push", "push", "peek", "pop", "peek"]
//[[], [1], [2], [], [], []]
// 输出：
//[null,null,null,1,null,2]
//
//
// 示例2:
//
//  输入：
//["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
//[[], [], [], [1], [], []]
// 输出：
//[null,null,null,null,null,true]
//
//
// 说明:
//
//
// 栈中的元素数目在[0, 5000]范围内。
//
//
// Related Topics 栈 设计 单调栈 👍 80 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
type SortedStack struct {
	Stack []int
	Top   int
}

//func Constructor() SortedStack {
//	return SortedStack{
//		Stack: []int{},
//		Top:   0,
//	}
//}

func (this *SortedStack) Push(val int) {
	if this.IsEmpty() {
		this.Stack = append(this.Stack, val)
		this.Top++
		return
	}
	if this.Stack[this.Top-1] < val {
		tmp := this.Peek()
		this.Pop()
		this.Push(val)
		this.Push(tmp)
	} else {
		this.Stack = append(this.Stack, val)
		this.Top++
	}
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.Stack = this.Stack[:this.Top-1]
	this.Top--
}

func (this *SortedStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Stack[this.Top-1]
}

func (this *SortedStack) IsEmpty() bool {
	return this.Top == 0
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
//leetcode submit region end(Prohibit modification and deletion)
