//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
// 示例 1：
//
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//
//
// 示例 2：
//
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//
//
// 提示：
//
//
// 2 <= timePoints <= 2 * 10⁴
// timePoints[i] 格式为 "HH:MM"
//
//
//
//
//
// 注意：本题与主站 539 题相同： https://leetcode-cn.com/problems/minimum-time-difference/
//
// Related Topics 数组 数学 字符串 排序 👍 34 👎 0

package main

import (
	"sort"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findMinDifference(timePoints []string) int {
	mins := make([]int, len(timePoints))
	for i := range timePoints {
		a := strings.Split(timePoints[i], ":")
		h, _ := strconv.Atoi(a[0])
		m, _ := strconv.Atoi(a[1])
		mins[i] = 60*h + m
	}
	sort.Ints(mins)
	minC := 24 * 60
	for i := 1; i < len(mins); i++ {
		if minC > mins[i]-mins[i-1] {
			minC = mins[i] - mins[i-1]
		}
	}
	if minC > 60*24-mins[len(mins)-1]+mins[0] {
		minC = 60*24 - mins[len(mins)-1] + mins[0]
	}
	return minC
}

//leetcode submit region end(Prohibit modification and deletion)
