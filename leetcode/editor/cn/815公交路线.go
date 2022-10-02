//给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。
//
//
// 例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1
//-> ... 这样的车站路线行驶。
//
//
// 现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。
//
// 求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
//输出：2
//解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
//
//
// 示例 2：
//
//
//输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= routes.length <= 500.
// 1 <= routes[i].length <= 10⁵
// routes[i] 中的所有值 互不相同
// sum(routes[i].length) <= 10⁵
// 0 <= routes[i][j] < 10⁶
// 0 <= source, target < 10⁶
//
//
// Related Topics 广度优先搜索 数组 哈希表 👍 316 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stationsMap := make(map[int][]int, 0)
	n := len(routes)
	edge := make([][]bool, n)
	for i, _ := range edge {
		edge[i] = make([]bool, n)
	}
	for i, _ := range routes {
		for j, _ := range routes[i] {
			if _, ok := stationsMap[routes[i][j]]; !ok {
				stationsMap[routes[i][j]] = make([]int, 0)
			}
			for _, v := range stationsMap[routes[i][j]] {
				edge[i][v] = true
				edge[v][i] = true
			}
			stationsMap[routes[i][j]] = append(stationsMap[routes[i][j]], i)
		}
	}
	starts := stationsMap[source]
	ends := stationsMap[target]
	distance := make([]int, n)
	for i, _ := range distance {
		distance[i] = -1
	}
	q := make([]int, 0)
	start := 0
	end := 0
	for _, line := range starts {
		distance[line] = 1
		q = append(q, line)
		end++
	}
	for start < end {
		line := q[start]
		start++
		for i, v := range edge[line] {
			if v && distance[i] == -1 {
				distance[i] = distance[line] + 1
				q = append(q, i)
				end++
			}
		}
	}
	min := int(^uint(0) >> 1)
	for _, v := range ends {
		if distance[v] != -1 && distance[v] < min {
			min = distance[v]
		}
	}
	if min == int(^uint(0)>>1) {
		return -1
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numBusesToDestination([][]int{
		{1, 7},
		{3, 6},
	}, 5, 5))
}
