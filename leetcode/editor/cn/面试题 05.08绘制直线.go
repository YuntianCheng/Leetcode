//已知一个由像素点组成的单色屏幕，每行均有 w 个像素点，所有像素点初始为 0，左上角位置为 (0,0)。
//
// 现将每行的像素点按照「每 32 个像素点」为一组存放在一个 int 中，再依次存入长度为 length 的一维数组中。
//
// 我们将在屏幕上绘制一条从点 (x1,y) 到点 (x2,y) 的直线（即像素点修改为 1），请返回绘制过后的数组。
//
//
//
// 注意：
//
//
// 用例保证屏幕宽度 w 可被 32 整除（即一个 int 不会分布在两行上）
//
//
//
//
// 示例1:
//
//
// 输入：length = 1, w = 32, x1 = 30, x2 = 31, y = 0
// 输出：[3]
// 解释：在第 0 行的第 30 位到第 31 位画一条直线，屏幕二进制形式表示为 [00000000000000000000000000000011]，因此
//返回 [3]
//
//
// 示例2:
//
//
// 输入：length = 3, w = 96, x1 = 0, x2 = 95, y = 0
// 输出：[-1, -1, -1]
// 解释：由于二进制 11111111111111111111111111111111 的 int 类型代表 -1，因此返回 [-1,-1,-1]
//
//
//
// 提示：
//
//
// 1 <= length <= 10^5
// 1 <= w <= 3 * 10^5
// 0 <= x1 <= x2 < w
// 0 <= y <= 10
//
//
// Related Topics 位运算 数组 数学 👍 24 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	array := make([]int, length)
	startIndex := (x1 + y*w) / 32
	endIndex := (x2 + y*w) / 32
	startBit := x1 % 32
	endBit := x2 % 32
	if startIndex != endIndex {
		array[startIndex] = int(int32(^uint32(0) >> startBit))
		for i := startIndex + 1; i < endIndex; i++ {
			array[i] = ^0
		}
		array[endIndex] = int(int32(^uint32(0) << (32 - (endBit + 1))))
	} else {
		tmp := ^uint32(0) >> startBit
		tmp = tmp & (^uint32(0) << (32 - (endBit + 1)))
		array[startIndex] = int(int32(tmp))
	}
	return array
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(drawLine(24, 96, 2, 19, 5))
	//fmt.Println(^int32(0))
}
