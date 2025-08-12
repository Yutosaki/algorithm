package main

import "fmt"

func main() {
	N, L, R := 0, 0, 0
	fmt.Scan(&N, &L, &R)
	s := ""
	fmt.Scan(&s)
	for i := L - 1; i < R; i++ {
		if s[i] != 'o' {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
