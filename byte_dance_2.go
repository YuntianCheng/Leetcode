package main

import "fmt"

func A2(n int) int64 {
	top := 1
	bottom := 1
	m := 2
	if n < 4 {
		m = n - 2
	}
	for i := 0; i < m; i++ {
		top = top * (n - i)
		bottom = bottom * (i + 1)
	}
	return int64(top / bottom)
}

func getPlanNum(distance int, coordinate []int) int {
	if len(coordinate) < 3 {
		return 0
	}
	var count int64 = 0
	start := 0
	end := 2
	for start <= len(coordinate)-3 {
		if start == end {
			end = end + 2
			continue
		}
		if coordinate[end]-coordinate[start] <= distance {
			if end < len(coordinate)-1 {
				end++
			} else {
				num := end - start + 1
				if num >= 3 {
					count += A2(num - 1)
					start++
				} else {
					break
				}
			}
		} else {
			num := end - start
			if num >= 3 {
				count += A2(num - 1)
			}
			start++
		}
	}
	return int(count % 99997867)
}

func main() {
	var N int
	var D int
	_, err := fmt.Scanln(&N, &D)
	if err != nil {
		return
	}
	coordinate := make([]int, N)
	coodinatePoints := make([]interface{}, N)
	for i := 0; i < N; i++ {
		coodinatePoints[i] = &coordinate[i]
	}
	_, err = fmt.Scanln(coodinatePoints...)
	if err != nil {
		return
	}
	fmt.Println(getPlanNum(D, coordinate))
}
