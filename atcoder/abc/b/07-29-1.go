package main

import "fmt"

func main() {
	S := ""
	fmt.Scan(&S)

	count := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'A' {
			for j := i; j < len(S); j++ {
				if S[j] == 'B' {
					diff := j - i
					if j+diff >= len(S) {
						break
					}
					if S[j+diff] == 'C' {
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)
}
