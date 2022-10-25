//给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
//
//
//
// 示例 1：
//
//
//输入:nums = [1,1,1], k = 2
//输出: 2
//解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
//
//
// 示例 2：
//
//
//输入:nums = [1,2,3], k = 3
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
//
//
//
// 注意：本题与主站 560 题相同： https://leetcode-cn.com/problems/subarray-sum-equals-k/
//
// Related Topics 数组 哈希表 前缀和 👍 110 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	var count, sum int
	m := make(map[int]int, 0)
	m[0] = 1
	for _, num := range nums {
		sum += num
		count += m[sum-k]
		m[sum]++
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(subarraySum([]int{1}, 0))
}
