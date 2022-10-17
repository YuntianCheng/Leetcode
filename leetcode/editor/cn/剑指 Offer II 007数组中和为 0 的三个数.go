//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//
//
// 示例 3：
//
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
//
// 注意：本题与主站 15 题相同：leetcode.cn/problems/3sum/
//
// Related Topics 数组 双指针 排序 👍 88 👎 0

package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; nums[i] <= 0 && i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			var j, k = i + 1, len(nums) - 1
			target := -nums[i]
			for j < k {
				sum := nums[j] + nums[k]
				if target < sum {
					k--
				} else if target == sum {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					m, n := j+1, k-1
					for m < k && nums[m] == nums[j] {
						m++
					}
					for n > j && nums[n] == nums[k] {
						n--
					}
					j, k = m, n
				} else {
					j++
				}
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}
