//堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行
//为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与
//普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行
//pop操作。
//
// 当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.
//
// 示例1:
//
//  输入：
//["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
//[[1], [1], [2], [1], [], []]
// 输出：
//[null, null, null, 2, 1, -1]
//
//
// 示例2:
//
//  输入：
//["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
//[[2], [1], [2], [3], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, 3]
//
//
// Related Topics 栈 设计 链表 👍 49 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)

type stackNode struct {
	top    int
	stack  []int
	next   *stackNode
	before *stackNode
}

type StackOfPlates struct {
	cap    int
	length int
	head   *stackNode
	end    *stackNode
}

//func Constructor(cap int) StackOfPlates {
//	stack := new(StackOfPlates)
//	stack.Init(cap)
//	return *stack
//}

func (stack *StackOfPlates) Init(cap int) {
	stack.cap = cap
	stack.length = 0
	stack.head = nil
	stack.end = nil
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	if this.length == 0 {
		this.head = &stackNode{
			top:    0,
			stack:  []int{},
			next:   nil,
			before: nil,
		}
		this.end = this.head
		this.length++
	}
	if this.end.top == this.cap {
		newStack := &stackNode{
			0, []int{}, nil, this.end,
		}
		this.end.next = newStack
		this.end = this.end.next
		this.length++
	}
	this.end.stack = append(this.end.stack, val)
	this.end.top++
}

func (this *StackOfPlates) Pop() int {
	if this.length == 0 {
		return -1
	}
	val := this.end.stack[this.end.top-1]
	this.end.top--
	if this.end.top == 0 {
		if this.length > 1 {
			this.end = this.end.before
			this.end.next = nil
		} else {
			this.head, this.end = nil, nil
		}
		this.length--
	} else {
		this.end.stack = this.end.stack[:this.end.top]
	}
	return val
}

func (this *StackOfPlates) PopAt(index int) int {
	if index >= this.length {
		return -1
	}
	var stack *stackNode
	if index > this.length/2 {
		index = this.length - index - 1
		stack = this.end
		for i := 0; i < index; i++ {
			stack = stack.before
		}
	} else {
		stack = this.head
		for i := 0; i < index; i++ {
			stack = stack.next
		}
	}
	val := stack.stack[stack.top-1]
	stack.top--
	if stack.top == 0 {
		if stack.before != nil {
			stack.before.next = stack.next
		} else {
			this.head = stack.next
		}
		if stack.next != nil {
			stack.next.before = stack.before
		} else {
			this.end = stack.before
		}
		this.length--
	} else {
		stack.stack = stack.stack[:stack.top]
	}
	return val
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := Constructor(2)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.PopAt(0)
	s.PopAt(0)
	s.PopAt(0)
}
