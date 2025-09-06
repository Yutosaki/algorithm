package main

import (
	"fmt"
)

func main() {
	n, l := 0, 0
	fmt.Scan(&n, &l)

	if l%3 != 0 {
		fmt.Println("0")
		return
	}

	d := make([]int, n)
	points := make(map[int]int)
	d[0] = 0
	points[0]++

	for i := 1; i <= n-1; i++ {
		fmt.Scan(&d[i])
		d[i] += d[i-1]
		points[d[i]%l]++
	}

	sum := 0
	for i := 0; i < l/3; i++ {
		sum += points[i] * points[(i+l/3)%l] * points[(i+2*l/3)%l]
	}

	fmt.Println(sum)
}
