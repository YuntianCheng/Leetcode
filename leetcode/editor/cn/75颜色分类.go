//给定一个包含红色、白色和蓝色、共 n 个元素的数组
// nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
//
// 必须在不使用库的sort函数的情况下解决这个问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
//
//
// 示例 2：
//
//
//输入：nums = [2,0,1]
//输出：[0,1,2]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//
//
//
//
// 进阶：
//
//
// 你可以不使用代码库中的排序函数来解决这道题吗？
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
//
// Related Topics 数组 双指针 排序 👍 1394 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	var red, white, blue int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			nums[blue] = 2
			nums[white] = 1
			nums[red] = 0
			blue++
			white++
			red++
		case 1:
			nums[blue] = 2
			nums[white] = 1
			blue++
			white++
		case 2:
			nums[blue] = 2
			blue++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
