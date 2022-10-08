//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
// 示例 1:
//
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//
// 示例 2：
//
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//
// 提示：
//
//
// 1 <= heights.length <=10⁵
// 0 <= heights[i] <= 10⁴
//
//
// Related Topics 栈 数组 单调栈 👍 2191 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	measure1 := make([]int, len(heights))
	measure2 := make([]int, len(heights))
	var singleStack = []int{0}
	measure1[0] = heights[0]
	for i := 1; i < len(heights); i++ {
		for len(singleStack) > 0 && heights[singleStack[len(singleStack)-1]] >= heights[i] {
			singleStack = singleStack[:len(singleStack)-1]
		}
		var count int
		if len(singleStack) == 0 {
			count = i + 1
		} else {
			count = i - singleStack[len(singleStack)-1]
		}
		measure1[i] = heights[i] * count
		singleStack = append(singleStack, i)
	}
	measure2[len(heights)-1] = heights[len(heights)-1]
	singleStack = []int{len(heights) - 1}
	for i := len(heights) - 2; i >= 0; i-- {
		for len(singleStack) > 0 && heights[singleStack[len(singleStack)-1]] >= heights[i] {
			singleStack = singleStack[:len(singleStack)-1]
		}
		var count int
		if len(singleStack) == 0 {
			count = len(heights) - i
		} else {
			count = singleStack[len(singleStack)-1] - i
		}
		measure2[i] = heights[i] * count
		singleStack = append(singleStack, i)
	}
	max := measure1[0] + measure2[0] - heights[0]
	for i := 1; i < len(heights); i++ {
		if measure1[i]+measure2[i]-heights[i] > max {
			max = measure1[i] + measure2[i] - heights[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5}))
}
