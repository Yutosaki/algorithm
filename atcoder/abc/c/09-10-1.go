package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type nums struct {
	count      int
	xCharacter int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	q := readInt()

	array := make([]nums, 0)
	for i := 0; i < q; i++ {
		queryType := readInt()
		if queryType == 1 {
			c := readInt()
			x := readInt()
			var newOne nums
			newOne.count = c
			newOne.xCharacter = x
			array = append(array, newOne)
		} else {
			k := readInt()

			sum := 0
			if len(array) > 0 {
				for i := 0; k-array[i].count > 0; {
					sum += array[i].count * array[i].xCharacter
					k -= array[i].count
					array = array[1:]
				}
				sum += k * array[0].xCharacter
				array[0].count -= k
				fmt.Println(sum)
			}
		}
	}
}
