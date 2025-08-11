package main

import (
	"fmt"
)

func main() {
	S := ""
	fmt.Scan(&S)

	var res float64
	for i := 0; i < len(S)-2; i++ {
		for j := i + 2; j < len(S); j++ {
			if S[i] == 't' && S[j] == 't' {
				t_len := j - i + 1
				x := 0
				for k := i; k <= j; k++ {
					if S[k] == 't' {
						x++
					}
				}
				tmp := (float64(x-2) / float64(t_len-2))

				if res < tmp {
					res = tmp
				}
			}
		}
	}

	fmt.Println(res)
}
