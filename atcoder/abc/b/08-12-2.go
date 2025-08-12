package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)

	flag := true
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			fmt.Print("#")
			flag = true
		} else if s[i] == '.' && flag {
			fmt.Print("o")
			flag = false
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
}
