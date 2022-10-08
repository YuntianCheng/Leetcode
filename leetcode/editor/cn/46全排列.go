//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2243 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func dfs10(nums []int, result *[][]int, path []int) {
	if len(nums) == 0 {
		*result = append(*result, append([]int{}, path...))
		return
	}
	for i, _ := range nums {
		path = append(path, nums[i])
		partOne := append([]int{}, nums[:i]...)
		partTwo := append([]int{}, nums[i+1:]...)
		dfs10(append(partOne, partTwo...), result, path)
		path = path[:len(path)-1]

	}
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfs10(nums, &result, path)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
