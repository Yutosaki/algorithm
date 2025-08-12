package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	x := 0
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if x == a[i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
