//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
// 说明：解集不能包含重复的子集。
//
// 示例:
//
//  输入： nums = [1,2,3]
// 输出：
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//
//
// Related Topics 位运算 数组 回溯 👍 98 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func dfs0804(nums []int, k int, result *[][]int) {
	if k == len(nums) {
		return
	}
	n := len(*result)
	for i := 0; i < n; i++ {
		tmp := append([]int{}, (*result)[i]...)
		*result = append(*result, append(tmp, nums[k]))
	}
	dfs0804(nums, k+1, result)
}
func subsets(nums []int) [][]int {
	result := make([][]int, 1, 1<<len(nums))
	result[0] = []int{}
	dfs0804(nums, 0, &result)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
