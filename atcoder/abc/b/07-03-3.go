package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	res := make([]int, 0)
	res = append(res, a[0])
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			res = append(res, a[i])
		}
	}
	fmt.Println(len(res))
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Println("")
}
