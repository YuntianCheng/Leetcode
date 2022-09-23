//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
//
// 示例 2:
//
//
//输入: numRows = 1
//输出: [[1]]
//
//
//
//
// 提示:
//
//
// 1 <= numRows <= 30
//
//
// Related Topics 数组 动态规划 👍 838 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	line := []int{0, 1, 0}
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = line[1 : len(line)-1]
		newLine := make([]int, 1, len(line)+2)
		newLine[0] = 0
		for j := 0; j < len(line)-1; j++ {
			newLine = append(newLine, line[j]+line[j+1])
		}
		newLine = append(newLine, 0)
		line = newLine
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
