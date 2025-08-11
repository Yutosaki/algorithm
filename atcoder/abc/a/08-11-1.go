package main

import "fmt"

func main() {
	N := 0
	A, B := 0, 0
	fmt.Scan(&N, &A, &B)
	S := ""
	fmt.Scan(&S)

	fmt.Println(S[A : N-B])
}
