//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Recortar el slice para darle longitud cero.
	s = s[:0]
	imprimirSlice(s)

	// Extender su longitud.
	s = s[:4]
	imprimirSlice(s)

	// Descartar los primeros dos valores.
	s = s[2:]
	imprimirSlice(s)
}

func imprimirSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
