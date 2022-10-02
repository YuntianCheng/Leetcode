//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 双指针 排序 👍 1256 👎 0

package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	difference := int(^uint(0) >> 1)
	var result int
	for start := 0; start <= len(nums)-3; start++ {
		end := len(nums) - 1
		for mid := start + 1; mid < end; {
			sum := nums[start] + nums[mid] + nums[end]
			var tmp int
			if target > sum {
				tmp = target - sum
				mid++
			} else {
				tmp = sum - target
				end--
			}
			if tmp == 0 {
				return sum
			}
			if difference > tmp {
				difference = tmp
				result = sum
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 2))
}
