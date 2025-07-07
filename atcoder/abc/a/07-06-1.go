package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	res := 0
	for i := 0; i < n; i++ {
		tmp:=0
		fmt.Scan(&tmp)
		if i%2 == 0 {
			res += tmp
		}
	}
	fmt.Println(res)
}
