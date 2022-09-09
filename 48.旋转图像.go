/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package main

// @lc code=start
func rotate2(matrix [][]int) {
	n := len(matrix)
	offset := (float64(n) - 1) / 2.0
	for i := 0.0; i < float64(len(matrix)); i++ {
		for j := 0.0; j < float64(len(matrix[int(i)])); j++ {
			if matrix[int(i)][int(j)] <= 1000 {
				x := j - offset
				y := offset - i
				temp := -x
				x = y
				y = temp
				k := offset - y
				l := x + offset
				nV := matrix[int(k)][int(l)]
				matrix[int(k)][int(l)] = matrix[int(i)][int(j)] + 2001
				for k != i || l != j {
					x = l - offset
					y = offset - k
					newY := -x
					x = y
					y = newY
					k = offset - y
					l = x + offset
					old := matrix[int(k)][int(l)]
					matrix[int(k)][int(l)] = nV + 2001
					nV = old
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = matrix[i][j] - 2001
		}
	}
}

// @lc code=end
