//n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。
//
// 如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。
//
// 给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。
//
//
//
//
// 示例 1：
//
//
//输入：stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
//输出：5
//解释：一种移除 5 块石头的方法如下所示：
//1. 移除石头 [2,2] ，因为它和 [2,1] 同行。
//2. 移除石头 [2,1] ，因为它和 [0,1] 同列。
//3. 移除石头 [1,2] ，因为它和 [1,0] 同行。
//4. 移除石头 [1,0] ，因为它和 [0,0] 同列。
//5. 移除石头 [0,1] ，因为它和 [0,0] 同行。
//石头 [0,0] 不能移除，因为它没有与另一块石头同行/列。
//
// 示例 2：
//
//
//输入：stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
//输出：3
//解释：一种移除 3 块石头的方法如下所示：
//1. 移除石头 [2,2] ，因为它和 [2,0] 同行。
//2. 移除石头 [2,0] ，因为它和 [0,0] 同列。
//3. 移除石头 [0,2] ，因为它和 [0,0] 同行。
//石头 [0,0] 和 [1,1] 不能移除，因为它们没有与另一块石头同行/列。
//
// 示例 3：
//
//
//输入：stones = [[0,0]]
//输出：0
//解释：[0,0] 是平面上唯一一块石头，所以不可以移除它。
//
//
//
// 提示：
//
//
// 1 <= stones.length <= 1000
// 0 <= xi, yi <= 10⁴
// 不会有两块石头放在同一个坐标点上
//
//
// Related Topics 深度优先搜索 并查集 图 👍 287 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
type UnionSet struct {
	count  int
	rank   map[int]int
	parent map[int]int
}

func (u *UnionSet) Count() int {
	return u.count
}

func (u *UnionSet) New() {
	u.count = 0
	u.parent = make(map[int]int, 0)
	u.rank = make(map[int]int, 0)
}

func (u *UnionSet) Find(x int) int {
	if p, ok := u.parent[x]; ok {
		if u.parent[p] != p {
			u.parent[x] = u.Find(p)
		}
	} else {
		u.parent[x] = x
		u.count++
	}
	return u.parent[x]
}

func (u *UnionSet) Merge(x, y int) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return
	}
	var rankX, rankY int
	if r, ok := u.rank[rootX]; ok {
		rankX = r
	} else {
		u.rank[rootX] = rankX
	}
	if r, ok := u.rank[rootY]; ok {
		rankY = r
	} else {
		u.rank[rootY] = rankY
	}
	if rankX > rankY {
		u.parent[rootY] = rootX
	} else {
		u.parent[rootX] = rootY
		if rankX == rankY {
			u.rank[rootY]++
		}
	}
	u.count--
}

func removeStones(stones [][]int) int {
	var u UnionSet
	u.New()
	for _, stone := range stones {
		u.Merge(stone[0]+10001, stone[1])
	}
	return len(stones) - u.Count()
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(removeStones([][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
		{2, 2},
	}))
}
