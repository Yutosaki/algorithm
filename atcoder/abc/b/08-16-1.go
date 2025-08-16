package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)

	taken := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&taken[i])
	}

	s := make(map[string]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				s[taken[i]+taken[j]]++
			}
		}
	}
	fmt.Println(len(s))
}
