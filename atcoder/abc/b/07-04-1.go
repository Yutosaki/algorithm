package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	s := make([][]rune, n)
	t := make([][]rune, n)

	for i := 0; i < n; i++ {
		text := ""
		fmt.Scan(&text)
		s[i] = []rune(text)
	}

	for i := 0; i < n; i++ {
		text := ""
		fmt.Scan(&text)
		t[i] = []rune(text)
	}

	min := 10001
	for rotate := 0; rotate < 4; rotate++ {
		counter := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if s[i][j] != t[i][j] {
					counter++
				}
			}
		}
		if counter+rotate < min {
			min = counter + rotate
		}

		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				s[i][j], s[j][i] = s[j][i], s[i][j]
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				s[i][j], s[i][n-j-1] = s[i][n-j-1], s[i][j]
			}
		}
	}

	fmt.Println(min)
}
