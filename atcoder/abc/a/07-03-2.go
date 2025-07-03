package main

import "fmt"

func main() {
	n, s := 0, 0
	fmt.Scan(&n, &s)

	t := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&t[i])
		if t[i]-t[i-1] > s {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
