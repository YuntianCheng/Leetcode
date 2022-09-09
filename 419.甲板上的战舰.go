/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// @lc code=start
func countBattleships(board [][]byte) int {
	mark := ""
	row := false
	count := 0
	for _, bytes := range board {
		tempMark := ""
		for i, b := range bytes {
			if !row {
				if b == 'X' {
					if i != len(bytes)-1 && bytes[i+1] == 'X' {
						row = true
						continue
					}
					tempMark += "'" + strconv.Itoa(i) + "'"
				}
			} else {
				if b != 'X' || i == len(bytes)-1 {
					row = false
					count++
				}
			}
			if b != 'X' && strings.Contains(mark, "'"+strconv.Itoa(i)+"'") {
				count++
				continue
			}
		}
		mark = tempMark
	}
	if len(mark) > 0 {
		strings.Trim(mark, "'")
		m := strings.Split(mark, "''")
		count += len(m)
	}
	return count
}

// @lc code=end
func main() {
	fmt.Println(countBattleships([][]byte{{'.'}}))
}
