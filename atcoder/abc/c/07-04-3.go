package main

import (
	"fmt"
)

type DSU struct {
	parentOrSize []int
}

func NewDSU(n int) *DSU {
	var d DSU
	d.parentOrSize = make([]int, n+1)

	for i := 1; i <= n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

func (d DSU) Find(index int) int {
	if d.parentOrSize[index] < 0 {
		return index
	}
	d.parentOrSize[index] = d.Find(d.parentOrSize[index])
	return d.parentOrSize[index]
}

func (d DSU) Merge(a, b int) int {
	x, y := d.Find(a), d.Find(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[a] < -d.parentOrSize[b] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Find(a)]
}

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	uf := NewDSU(n)

	degree := make([]int, n+1)

	for i := 0; i < m; i++ {
		a, b := 0, 0
		fmt.Scan(&a, &b)
		degree[a]++
		degree[b]++
		uf.Merge(a, b)
	}

	for i := 1; i <= n; i++ {
		if degree[i] != 2 {
			fmt.Println("No")
			return
		}
	}

	if uf.Size(1) != n {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
