//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 测试用例的答案是一个 32-位 整数。
//
// 子数组 是数组的连续子序列。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
//
//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -10 <= nums[i] <= 10
// nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
//
//
// Related Topics 数组 动态规划 👍 1783 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct1(nums []int) int {

	result := ^int(^uint(0) >> 1)
	min := 1
	max := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := max
			max = min
			min = tmp
		}
		if nums[i] > nums[i]*max {
			max = nums[i]
		} else {
			max *= nums[i]
		}
		if nums[i] < nums[i]*min {
			min = nums[i]
		} else {
			min *= nums[i]
		}
		if result < max {
			result = max
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
