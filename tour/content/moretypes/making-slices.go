//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	a := make([]int, 5)
	imprimirSlice("a", a)

	b := make([]int, 0, 5)
	imprimirSlice("b", b)

	c := b[:2]
	imprimirSlice("c", c)

	d := c[2:5]
	imprimirSlice("d", d)
}

func imprimirSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
