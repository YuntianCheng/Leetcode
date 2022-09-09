package main

import (
	"fmt"
)

//BALLOON
func pre(blocks []int, index int) int {
	num := index
	for i := index - 1; i >= 0; i-- {
		if blocks[i] >= blocks[num] {
			num = i
			continue
		}
		break
	}
	return num
}
func after(blocks []int, index int) int {
	num := index
	for i := index + 1; i < len(blocks); i++ {
		if blocks[i] >= blocks[num] {
			num = i
			continue
		}
		break
	}
	return num
}
func Solution2(blocks []int) int {
	// write your code in Go 1.4
	max := 0
	for i := 0; i < len(blocks); {
		left := pre(blocks, i)
		right := after(blocks, i)
		if right-left+1 > max {
			max = right - left + 1
		}
		j := right + 1
		for j+1 < len(blocks) {
			if blocks[j] > blocks[j+1] {
				j++
				continue
			}
			break
		}
		i = j
	}
	return max
}

func main() {
	fmt.Println(Solution2([]int{1, 1}))
}
