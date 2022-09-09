package main

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	var cMap map[string]int = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	num := 0
	length := len(s)
	for i := 0; i < length; i++ {
		if i == length-1 {
			num += cMap[string(s[i])]
			return num
		}
		if cMap[string(s[i])] < cMap[string(s[i+1])] {
			num += cMap[string(s[i])+string(s[i+1])]
			i++
			continue
		}
		num += cMap[string(s[i])]
	}
	return num
}

// @lc code=end
