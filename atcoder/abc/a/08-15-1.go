package main

import "fmt"

func main() {
	n, m := 0, 0

	fmt.Scan(&n, &m)

	counter := 0
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		counter += tmp
	}
	if counter <= m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
