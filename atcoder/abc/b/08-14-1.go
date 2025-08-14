package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)

	flag := false
	before := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "#" {
			if flag {
				fmt.Printf("%d,%d\n", before+1, i+1)
				flag = false
			} else {
				before = i
				flag = true
			}
		}
	}
}
