package main

import (
	"fmt"
	"sort"
)

func main() {
	case_num := 0
	fmt.Scan(&case_num)

	for i := 0; i < case_num; i++ {
		n := 0
		fmt.Scan(&n)
		s := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&s[j])
		}

		sort.Ints(s[1 : n-1])

		current := 0
		counter := 1

		for 2*s[current] < s[n-1] {
			left, right := current, n-1

			for (right - left) > 1 {
				mid := left + (right-left)/2
				if 2*s[current] < s[mid] {
					right = mid
				} else {
					left = mid
				}
			}
			if 2*s[current] < s[left] {
				break
			} else if s[current] >= s[left] {
				break
			}
			counter++
			current = left
		}

		counter++
		if 2*s[current] >= s[n-1] {
			fmt.Println(counter)
		} else {
			fmt.Println(-1)
		}
	}
}
