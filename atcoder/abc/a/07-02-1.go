package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	S, T := "", ""
	fmt.Scan(&S, &T)
	for i := 0; i < n; i++ {
		if S[i] == 'o' && S[i] == T[i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
	return
}
