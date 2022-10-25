//给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
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
//
//
//
// 注意：本题与主站 84 题相同： https://leetcode-cn.com/problems/largest-rectangle-in-
//histogram/
//
// Related Topics 栈 数组 单调栈 👍 65 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	var stackLeft = make([]int, 0)
	var stackRight = make([]int, 0)
	var n = len(heights) - 1
	left, right := make([]int, len(heights)), make([]int, len(heights))
	var maxm int
	for i := range heights {
		for len(stackLeft) > 0 && heights[stackLeft[len(stackLeft)-1]] >= heights[i] {
			stackLeft = stackLeft[:len(stackLeft)-1]
		}
		for len(stackRight) > 0 && heights[stackRight[len(stackRight)-1]] >= heights[n-i] {
			stackRight = stackRight[:len(stackRight)-1]
		}
		var count int
		if len(stackLeft) == 0 {
			count = i + 1
		} else {
			count = i - stackLeft[len(stackLeft)-1]
		}
		left[i] = count * heights[i]
		stackLeft = append(stackLeft, i)
		if len(stackRight) == 0 {
			count = i + 1
		} else {
			count = stackRight[len(stackRight)-1] - (n - i)
		}
		right[n-i] = count * heights[n-i]
		stackRight = append(stackRight, n-i)
	}
	for i := range heights {
		if maxm < left[i]+right[i]-heights[i] {
			maxm = left[i] + right[i] - heights[i]
		}
	}
	return maxm
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
