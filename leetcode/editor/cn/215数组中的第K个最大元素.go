//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 1911 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left
		var i, j = left + 1, right
		for i <= j {
			if nums[i] > nums[mid] {
				i++
				continue
			}
			if nums[j] <= nums[mid] {
				j--
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[mid], nums[j] = nums[j], nums[mid]
		mid = j
		if mid == k-1 {
			return nums[mid]
		}
		if mid > k-1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findKthLargest([]int{2, 1}, 2))
}
