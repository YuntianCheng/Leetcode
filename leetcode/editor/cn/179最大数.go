//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
//
//输入：nums = [10,2]
//输出："210"
//
// 示例 2：
//
//
//输入：nums = [3,30,34,5,9]
//输出："9534330"
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 10⁹
//
//
// Related Topics 贪心 字符串 排序 👍 1029 👎 0

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	for i := range nums {
		str[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(str, func(i, j int) bool {
		if str[i][0] == str[j][0] {
			return str[i]+str[j] > str[j]+str[i]
		}
		return str[i] > str[j]
	})
	result := strings.Join(str, "")
	result = strings.TrimLeft(result, "0")
	if result == "" {
		return "0"
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(largestNumber([]int{0, 0}))
}
