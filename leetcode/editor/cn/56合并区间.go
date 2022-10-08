//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics 数组 排序 👍 1661 👎 0

package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//result := make([][]int, 0)
	var i int
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
			j++
		}
		if j > i+1 {
			partTwo := append([][]int{}, intervals[j:]...)
			intervals = append(intervals[:i], []int{start, end})
			intervals = append(intervals, partTwo...)
		}
		i++

	}
	return intervals
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(merge([][]int{
		{2, 3},
		{2, 2},
		{3, 3},
		{1, 3},
		{5, 7},
		{2, 2},
		{4, 6},
	}))
}
