//可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
//
// 'A'：Absent，缺勤
// 'L'：Late，迟到
// 'P'：Present，到场
//
//
// 如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//
// 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
// 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
//
//
// 给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。答案可能很大，所以返回对 10⁹ + 7
//取余 的结果。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：8
//解释：
//有 8 种长度为 2 的记录将被视为可奖励：
//"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//只有"AA"不会被视为可奖励，因为缺勤次数为 2 次（需要少于 2 次）。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：3
//
//
// 示例 3：
//
//
//输入：n = 10101
//输出：183236316
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
//
//
// Related Topics 动态规划 👍 279 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(n int) int {
	//对于某个长度为i的记录序列，以下变量表示对应情况的序列数量
	//P: 序列中没有A，最新的一个记录不是L
	//AP: 序列中有过A，最新的一个记录不是L
	//L: 序列中没有A，最新的一个记录是L
	//AL: 序列中有过A，最新的一个记录是L
	//LL: 序列中没有A，最新的两个记录是LL
	//ALL: 序列中有过A，最新的两个记录是LL
	//A: 最新的一个记录是A
	var P, L, LL, AP, AL, ALL, A = 1, 1, 0, 0, 0, 0, 1
	for i := 2; i <= n; i++ {
		P, L, LL, AP, AL, ALL, A = (P+L+LL)%1000000007, P, L, (AP+A+AL+ALL)%1000000007, (A+AP)%1000000007, AL, (P+L+LL)%1000000007
	}
	return (P + L + LL + AP + AL + ALL + A) % 1000000007
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(checkRecord(10101))
}
