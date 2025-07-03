package main

import (
	"fmt"
)

func main() {
	n, q := 0, 0
	fmt.Scan(&n, &q)

	a := make([]int, q)
	row := make([]int, n+2)

	counter := 0
	for i := 0; i < q; i++ {
		fmt.Scan(&a[i])
		if row[a[i]] == 0 {
			row[a[i]] = 1
			if row[a[i]-1] == 0 && row[a[i]+1] == 0 {
				counter++
			} else if row[a[i]-1] == 1 && row[a[i]+1] == 1 {
				counter--
			}
		} else {
			row[a[i]] = 0
			if row[a[i]-1] == 0 && row[a[i]+1] == 0 {
				counter--
			} else if row[a[i]-1] == 1 && row[a[i]+1] == 1 {
				counter++
			}
		}
		fmt.Println(counter)
	}
}
