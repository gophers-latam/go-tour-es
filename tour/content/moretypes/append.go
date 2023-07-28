//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	var s []int
	imprimirSlice(s)

	// append funciona en slices nil.
	s = append(s, 0)
	imprimirSlice(s)

	// El slice crece según sea necesario.
	s = append(s, 1)
	imprimirSlice(s)

	// Se pueden agregar más de un elemento a la vez.
	s = append(s, 2, 3, 4)
	imprimirSlice(s)
}

func imprimirSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
