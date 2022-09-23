//给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//解释：[4,9] 也是可通过的
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
//
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 624 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	numsMap := make(map[int]bool, len(nums1))
	for _, num := range nums1 {
		numsMap[num] = true
	}
	if len(nums1) > len(nums2) {
		result = make([]int, 0, len(nums2))
	} else {
		result = make([]int, 0, len(nums2))
	}
	for _, num := range nums2 {
		if val, ok := numsMap[num]; ok && val {
			result = append(result, num)
			numsMap[num] = false
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
