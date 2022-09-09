package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
var cMap map[int]string = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}
var roman = strings.Builder{}

func intToRoman(num int) string {

	if num >= 1000 {
		countM := num / 1000
		for i := 0; i < countM; i++ {
			roman.WriteString(cMap[1000])
		}
		num = num - countM*1000
	}
	if num >= 900 {
		roman.WriteString(cMap[900])
		num = num - 900
	}
	if num >= 500 {
		roman.WriteString(cMap[500])
		num = num - 500
	}
	if num >= 400 {
		roman.WriteString(cMap[400])
		num = num - 400
	}
	if num >= 100 {
		countC := num / 100
		for i := 0; i < countC; i++ {
			roman.WriteString(cMap[100])
		}
		num = num - countC*100
	}
	if num >= 90 {
		roman.WriteString(cMap[90])
		num = num - 90
	}
	if num >= 50 {
		roman.WriteString(cMap[50])
		num = num - 50
	}
	if num >= 40 {
		roman.WriteString(cMap[40])
		num = num - 40
	}
	if num >= 10 {
		countX := num / 10
		for i := 0; i < countX; i++ {
			roman.WriteString(cMap[10])
		}
		num = num - 10*countX
	}
	if num == 9 {
		roman.WriteString(cMap[9])
		num = num - 9
	}
	if num >= 5 {
		roman.WriteString(cMap[5])
		num = num - 5
	}
	if num == 4 {
		roman.WriteString(cMap[4])
		num = num - 4
	}
	if num >= 1 {
		for i := 0; i < num; i++ {
			roman.WriteString(cMap[1])
		}
	}
	return roman.String()
}

// @lc code=end
