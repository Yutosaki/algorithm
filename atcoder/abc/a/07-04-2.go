package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)

	seen := make([]bool, 26)

	for _, char := range s {
		seen[char-'a'] = true
	}

	for i := 0; i < 26; i++ {
		if !seen[i] {
			fmt.Println(string('a' + i))
			return
		}
	}
}
