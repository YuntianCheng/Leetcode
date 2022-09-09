/*
 * @lc app=leetcode.cn id=715 lang=golang
 *
 * [715] Range 模块
 */
package main

// @lc code=start
type Range struct {
	Left   int
	Right  int
	Before *Range
	Next   *Range
}
type RangeModule struct {
	RangeNum int
	Starts   []int
	Start    *Range
	End      *Range
}

func Constructor() RangeModule {
	return RangeModule{
		RangeNum: 0,
		Starts:   make([]int, 0),
		Start:    nil,
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	if this.RangeNum == 0 {
		this.Start = &Range{
			Left:  left,
			Right: right,
		}
		this.End = this.Start
		this.RangeNum++
		return
	}
	cursor1 := this.Start
	cursor2 := this.End
	for {
		if cursor1 == nil || cursor1.Left > right {
			break
		}
		cursor1 = cursor1.Next
	}
	for {
		if cursor2 == nil || cursor2.Right < left {
			break
		}
		cursor2 = cursor2.Before
	}
}

func (this *RangeModule) QueryRange(left int, right int) bool {

}

func (this *RangeModule) RemoveRange(left int, right int) {

}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
// @lc code=end
