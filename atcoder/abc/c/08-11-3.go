package main

import (
	"fmt"
)

func main() {
	N := 0
	fmt.Scan(&N)
	dp := make(map[int]int)
	counter := 0

	for i := 0; i < N; i++ {
		a := 0
		fmt.Scan(&a)

		target := i - a
		counter += dp[target]

		dp[i+a]++
	}
	fmt.Println(counter)
}
