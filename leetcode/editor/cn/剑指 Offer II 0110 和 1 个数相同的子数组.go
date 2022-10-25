//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
// 示例 1：
//
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
//
// 示例 2：
//
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 不是 0 就是 1
//
//
//
//
//
// 注意：本题与主站 525 题相同： https://leetcode-cn.com/problems/contiguous-array/
//
// Related Topics 数组 哈希表 前缀和 👍 104 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums []int) int {
	var sum, length int
	m := make(map[int]int, 0)
	m[0] = -1
	for j, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}
		if v, ok := m[sum]; ok {
			if length < j-v {
				length = j - v
			}
		} else {
			m[sum] = j
		}
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)
