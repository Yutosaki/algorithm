package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextString()

	srunes := []rune(s)

	has := true

	for i := 2; i < len(srunes); i++ {
		if srunes[i] >= 65 && srunes[i] <= 90 {
			has = strings.ContainsRune(t, srunes[i-1])
			if !has {
				break
			}
		}
	}

	if has {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
