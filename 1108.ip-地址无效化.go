/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */
package main

import "strings"

// @lc code=start
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// @lc code=end
