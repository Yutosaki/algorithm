package main

import (
	"fmt"
)

func main() {
	N, M := 0, 0
	fmt.Scan(&N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < M; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		for j := 0; j < N; j++ {
			if A[j] == tmp {
				A[j] = -1
				break
			}
		}
	}

	for i := 0; i < N; i++ {
		if A[i] != -1 {
			fmt.Print(A[i], " ")
		}
	}
	fmt.Println()
}
