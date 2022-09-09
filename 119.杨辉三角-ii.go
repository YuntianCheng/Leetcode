/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */
package main

import (
	"fmt"
	"math/big"
)

// @lc code=start
func getRow(rowIndex int) []int {
	var result []int
	for i := 0; i <= rowIndex; i++ {
		result = append(result, C(rowIndex, i))
	}
	return result
}

func C(n int, m int) int {
	if m == 0 {
		return 1
	}
	k := big.NewInt(1)
	l := big.NewInt(1)
	for i := n; i >= n-m+1; i-- {
		k = k.Mul(k, big.NewInt(int64(i)))
	}
	for i := m; i >= 1; i-- {
		l = l.Mul(l, big.NewInt(int64(i)))
	}
	r := k.Div(k, l).Int64()
	return int(r)
}

// @lc code=end
func main() {
	fmt.Println(getRow(3))
}
