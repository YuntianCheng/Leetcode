//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
//
//
//
// 示例 1:
//
// 输入: [7,5,6,4]
//输出: 5
//
//
//
// 限制：
//
// 0 <= 数组长度 <= 50000
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 874 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	var cnt int
	for i := 1; i < len(nums); i++ {
		var j = i
		for ; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
		cnt += j
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
