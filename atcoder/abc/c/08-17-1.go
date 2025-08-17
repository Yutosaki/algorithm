package main

import (
	"fmt"
)

func main() {
	t := 0
	fmt.Scan(&t)
	for t > 0 {
		t--

		n := 0
		s := ""
		fmt.Scan(&n, &s)

		s_zero := "0" + s
		ok := make([]bool, 1<<n)

		ok[0] = true
		for i := 0; i < (1 << n); i++ {
			if !ok[i] {
				continue
			}

			for j := 0; j < n; j++ {
				if i&(1<<j) == 0 {
					next := i | 1<<j
					if s_zero[next] == '0' {
						ok[next] = true
					}
				}
			}
		}

		if ok[(1<<n)-1] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
