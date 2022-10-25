//给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。
//
//
//
// 示例 1:
//
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
//
// 示例 2:
//
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
//
// 示例 3:
//
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//
//
// 提示:
//
//
// 1 <= nums1.length, nums2.length <= 10⁴
// -10⁹ <= nums1[i], nums2[i] <= 10⁹
// nums1, nums2 均为升序排列
// 1 <= k <= 1000
//
//
//
//
//
// 注意：本题与主站 373 题相同：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-
//sums/
//
// Related Topics 数组 堆（优先队列） 👍 51 👎 0

package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var makeHeap func([]int)
	var popHeap func() []int
	var heap = make([][]int, 0)
	makeHeap = func(value []int) {
		heap = append(heap, value)
		var i = len(heap) - 1
		for (i-1)/2 >= 0 {
			if heap[i][0]+heap[i][1] < heap[(i-1)/2][0]+heap[(i-1)/2][1] {
				heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
				i = (i - 1) / 2
			} else {
				break
			}
		}
	}
	popHeap = func() (result []int) {
		result = heap[0]
		heap[len(heap)-1], heap[0] = heap[0], heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		var i = 0
		for 2*i+1 < len(heap) {
			var l, r = -1000000001, -1000000001
			l = heap[2*i+1][0] + heap[2*i+1][1]
			if 2*i+2 < len(heap) {
				r = heap[2*i+2][0] + heap[2*i+2][1]
			}
			if l < r {
				heap[2*i+1], heap[i] = heap[i], heap[2*i+1]
				i = 2*i + 1
			} else if r > -1000000001 {
				heap[2*i+2], heap[i] = heap[i], heap[2*i+2]
				i = 2*i + 2
			} else {
				break
			}
		}
		return
	}
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			makeHeap([]int{n1, n2})
		}
	}
	var r = make([][]int, 0)
	for k > 0 && len(heap) > 0 {
		r = append(r, popHeap())
		k--
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}
