package main

import "fmt"

func main() {
	n, age := 0, 0

	fmt.Scan(&n)

	race_age := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&race_age[i])
	}
	fmt.Scan(&age)

	counter := 0
	for i := 0; i < n; i++ {
		if age <= race_age[i] {
			counter++
		}
	}
	fmt.Println(counter)
}
