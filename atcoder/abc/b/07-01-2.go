package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	ball := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&ball[i])
	}

	result := make([]int, n)
	order := make([]int, k)

	for i := 0; i < k; i++ {
		if ball[i] == 0 {
			min := result[0]
			tmp := 0
			for j := 0; j < n; j++ {
				if result[j] < min {
					min = result[j]
					tmp = j
				}
			}
			result[tmp]++
			order[i] = tmp + 1
		} else {
			result[ball[i]-1]++
			order[i] = ball[i]
		}
	}

	for i := 0; i < k; i++ {
		fmt.Print(order[i], " ")
	}
	fmt.Println("")
}
