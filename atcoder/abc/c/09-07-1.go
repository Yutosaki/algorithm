package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func readInt() int {
	n := 0
	fmt.Fscan(reader, &n)
	return n
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	n := readInt()
	q := readInt()

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i + 1
	}
	current := 0

	for i := 0; i < q; i++ {
		numType := 0
		fmt.Scan(&numType)
		if numType == 1 {
			p := readInt()
			x := readInt()
			array[(current+p-1)%n] = x
		} else if numType == 2 {
			p := readInt()
			fmt.Println(array[(current+p-1)%n])
		} else {
			k := readInt()
			current += k
			current %= n
		}
	}
}
