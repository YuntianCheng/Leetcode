package main

import (
	"fmt"
	"sort"
)

func binaryGetNum(nums []int, min, max, start, end int) int {
	var count int
	if start < end {
		mid := (start + end) / 2
		if nums[mid] < min {
			count += binaryGetNum(nums, min, max, mid+1, end)
		} else if nums[mid] > max {
			count += binaryGetNum(nums, min, max, start, mid-1)
		} else {
			count += binaryGetNum(nums, min, max, mid+1, end)
			count += binaryGetNum(nums, min, max, start, mid)
		}
	} else {
		if nums[start] >= min && nums[start] <= max {
			count++
		}
	}
	return count
}

func niasdassd(nums []int, min, max int) int {
	sort.Ints(nums)
	var count int
	for k := 0; k < len(nums)-1; k++ {
		l := min - nums[k]
		h := max - nums[k]
		start := k + 1
		end := len(nums) - 1
		count += binaryGetNum(nums, l, h, start, end)
	}
	return count
}

func getSK(a string, k int) string {
	pos := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			pos = append(pos, i)
		}
	}
	var i int
	minLength := int(^uint(0) >> 1)
	starts := make([]int, 0)
	for j := k - 1; j < len(pos); j++ {
		if minLength > pos[j]-pos[i]+1 {
			minLength = pos[j] - pos[i] + 1
			starts = starts[:0]
			starts = append(starts, pos[i])
		} else if minLength == pos[j]-pos[i]+1 {
			starts = append(starts, pos[i])
		}
		i++
	}
	result := a[starts[0] : starts[0]+k+1]
	for i = 1; i < len(starts); i++ {
		if result > a[starts[i]:starts[i]+k+1] {
			result = a[starts[i] : starts[i]+k+1]
		}
	}
	return result
}

func main() {
	fmt.Println(niasdassd([]int{2, 3, 4, 5}, 5, 7))
	fmt.Println(getSK("100101", 2))
}
