//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 1919 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0)
	var result []int
	for i := 0; i < k; i++ {
		var j = len(q) - 1
		for ; j >= 0; j-- {
			if nums[q[j]] > nums[i] {
				break
			}
		}
		q = q[:j+1]
		q = append(q, i)
	}
	result = append(result, nums[q[0]])
	for r := k; r < len(nums); r++ {
		l := r - k + 1
		if q[0] < l {
			q = q[1:]
		}
		var j = len(q) - 1
		for ; j >= 0; j-- {
			if nums[q[j]] > nums[r] {
				break
			}
		}
		q = q[:j+1]
		q = append(q, r)
		result = append(result, nums[q[0]])
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
