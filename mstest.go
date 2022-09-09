package main

import (
	"fmt"
)

//BALLOON
func Solution(S string) int {
	// write your code in Go 1.4
	var sta = [5]int{0, 0, 0, 0, 0}
	for i := 0; i < len(S); i++ {
		if S[i] == 'B' {
			sta[0]++
		} else if S[i] == 'A' {
			sta[1]++
		} else if S[i] == 'L' {
			sta[2]++
		} else if S[i] == 'O' {
			sta[3]++
		} else if S[i] == 'N' {
			sta[4]++
		}
	}
	sta[2] = sta[2] / 2
	sta[3] = sta[3] / 2
	min := sta[0]
	for _, s := range sta[1:] {
		if s < min {
			min = s
		}
	}
	return min
}

func main() {
	fmt.Println(Solution("BAOOLLNNOLOLGBAX"))
}
