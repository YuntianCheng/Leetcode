//节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
//
// 示例1:
//
//  输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
// 输出：true
//
//
// 示例2:
//
//  输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [
//1, 3], [2, 3], [3, 4]], start = 0, target = 4
// 输出 true
//
//
// 提示：
//
//
// 节点数量n在[0, 1e5]范围内。
// 节点编号大于等于 0 小于 n。
// 图中可能存在自环和平行边。
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 👍 68 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	charts := make([][]int, n)
	m := make([]bool, n)
	m[start] = true
	for i := range charts {
		charts[i] = make([]int, 0)
	}
	for i := range graph {
		x := graph[i][0]
		y := graph[i][1]
		charts[x] = append(charts[x], y)
	}
	queue := make([]int, 1)
	queue[0] = start
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, k := range charts[node] {
			if !m[k] {
				if k == target {
					return true
				}
				queue = append(queue, k)
				m[k] = true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findWhetherExistsPath(3, [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{1, 2},
	}, 0, 2))
}
