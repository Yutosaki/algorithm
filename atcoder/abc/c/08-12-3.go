package main

import (
	"fmt"
	"sort"
)

func main() {
	n, k, x := 0, 0, 0
	fmt.Scan(&n, &k, &x)
	s := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	all := make([]string, 0)
	for i := 0; i < n; i++ {
		if k > 1 {
			for j := 0; j < n; j++ {
				if k > 2 {
					for l := 0; l < n; l++ {
						if k > 3 {
							for p := 0; p < n; p++ {
								if k > 4 {
									for q := 0; q < n; q++ {
										all = append(all, s[i]+s[j]+s[l]+s[p]+s[q])
									}
								} else {
									all = append(all, s[i]+s[j]+s[l]+s[p])
								}
							}
						} else {
							all = append(all, s[i]+s[j]+s[l])
						}
					}
				} else {
					all = append(all, s[i]+s[j])
				}
			}
		} else {
			all = append(all, s[i])
		}
	}

	sort.Strings(all)
	fmt.Println(all[x-1])
}
