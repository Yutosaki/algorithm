package main

import "fmt"

func main() {
	n, l, r := 0, 0, 0
	fmt.Scan(&n, &l, &r)

	counter := 0
	for i := 0; i < n; i++ {
		start, end := 0, 0
		fmt.Scan(&start, &end)
		if start <= l && r <= end {
			counter++
		}
	}
	fmt.Println(counter)
}
